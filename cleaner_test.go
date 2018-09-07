package main

import (
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

	_, err := executeCommand(params...)

	assert.NotNil(t, err, "should return error if function is called without parameters")
}

func TestShouldReturnErrorNilIfThereAreParameters(t *testing.T) {
	params := []string{"env"}

	_, err := executeCommand(params...)

	assert.Nil(t, err, "should return nil if function has parameters")
}
