package main

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repository string
var tags []string

func TestShoulReturnErrorIfDoesNotdHaveARepo(t *testing.T) {
	repository = ""
	tags = []string{}

	err := CleanRegistry(repository, tags)

	assert.NotNil(t, err)
}

func TestShouldReturnErrorIfDoesNotHaveTagsAsArguments(t *testing.T) {
	repository = "repo"
	tags = []string{}

	err := CleanRegistry(repository, tags)

	assert.NotNil(t, err)
}

func TestShouldReturnTrueWhenValueExistsInArray(t *testing.T) {
	tag := "b-1"
	tags = []string{"b-1", "b-2"}

	result := inArray(tag, tags)

	assert.True(t, result, "Should return true if value is in array")
}
func TestShouldReturnFalseWhenValueDoesNotExistInArray(t *testing.T) {
	tag := "b-3"
	tags = []string{"b-1", "b-2"}

	result := inArray(tag, tags)

	assert.False(t, result, "Should return false if value is not in array")
}

func TestShouldReturnErrorIfThereAreNotParameters(t *testing.T) {
	params := []string{}

	_, err := ExecuteCommandFunction(params...)

	assert.NotNil(t, err, "should return error if function is called without parameters")
}

func TestShouldReturnErrorNilIfThereAreParameters(t *testing.T) {
	params := []string{"env"}

	_, err := ExecuteCommandFunction(params...)

	assert.Nil(t, err, "should return nil if function has parameters")
}

func TestShouldReturnErrorIfExecuteCommandFunctionGoesWrong(t *testing.T) {
	repository = "demo"
	tags = []string{"demo1", "demo2"}

	oldExecuteCommandFunction := ExecuteCommandFunction
	defer func() {
		ExecuteCommandFunction = oldExecuteCommandFunction
	}()
	ExecuteCommandFunction = func(params ...string) ([]byte, error) {
		return nil, errors.New("Something goes wrong executing command")
	}

	err := getTags(repository, tags)

	assert.NotNil(t, err, "should return error if execute command goes wrong")
}

func TestShouldReturnErrorIfOutIsBad(t *testing.T) {
	var response []string
	//s = make([]byte, 5, 5)

	err := UnmarshalFunction(nil, &response)

	assert.NotNil(t, err, "should return error if unmarshall goes wrong")
}

func TestShouldReturnNilIfOutIsGood(t *testing.T) {
	var response []string
	out := []string{
		"demo1",
	}
	bytes, _ := json.Marshal(out)

	err := UnmarshalFunction(bytes, &response)

	assert.Nil(t, err, "should return nil if unmarshall goes fine")
}

func TestShouldReturnErrorIfGetErrorFromUnmarshal(t *testing.T) {
	repository = "demo"
	tags = []string{
		"demo3",
	}

	oldExecuteCommandFunction := ExecuteCommandFunction
	oldUnmarshalFunction := UnmarshalFunction
	defer func() {
		ExecuteCommandFunction = oldExecuteCommandFunction
		UnmarshalFunction = oldUnmarshalFunction
	}()

	ExecuteCommandFunction = func(params ...string) ([]byte, error) {
		out := []string{
			"demo1",
			"demo2",
		}
		bytes, _ := json.Marshal(out)
		return bytes, nil
	}

	UnmarshalFunction = func(bytes []byte, response interface{}) error {
		return errors.New("something got wrong with json marshal")
	}

	err := getTags(repository, tags)

	assert.NotNil(t, err, "should return error if json unmarshal got wrong")
}

func TestShouldReturnNilWhenDeletedUnusedTagsWorks(t *testing.T) {
	repository = "demo"
	tags = []string{
		"demo3",
	}

	oldExecuteCommandFunction := ExecuteCommandFunction
	oldUnmarshalFunction := UnmarshalFunction
	oldDeleteUnusedTags := DeleteUnusedTags
	defer func() {
		ExecuteCommandFunction = oldExecuteCommandFunction
		UnmarshalFunction = oldUnmarshalFunction
		DeleteUnusedTags = oldDeleteUnusedTags
	}()

	ExecuteCommandFunction = func(params ...string) ([]byte, error) {
		out := []string{
			"demo1",
			"demo2",
		}
		bytes, _ := json.Marshal(out)
		return bytes, nil
	}

	UnmarshalFunction = func(bytes []byte, response interface{}) error {
		return nil
	}

	DeleteUnusedTags = func(tag string, repository string) {

	}

	err := getTags(repository, tags)

	assert.Nil(t, err, "should return nil when got fine")
}
