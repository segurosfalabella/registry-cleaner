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
	var tag string = "b-1"
	tags = []string{"b-1", "b-2"}

	result := in_array(tag, tags)

	assert.True(t, result, "should return true if value is in array")
}
func TestShouldReturnFalseWhenValueDoesNotExistInArray(t *testing.T) {
	var tag string = "b-3"
	tags = []string{"b-1", "b-2"}

	result := in_array(tag, tags)

	assert.False(t, result, "should return false if value is not in array")
}
