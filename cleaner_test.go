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
