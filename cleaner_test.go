package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var repo string
var tags []string

func TestShoulReturnErrorIfDoesNotdHaveARepo(t *testing.T) {
	repo = ""
	tags = []string{}

	err := CleanRegistry(repo, tags)

	assert.NotNil(t, err)
}

func TestShouldReturnErrorIfDoesNotHaveTagsAsArguments(t *testing.T) {
	repo = "repo"
	tags = []string{}

	err := CleanRegistry(repo, tags)

	assert.NotNil(t, err)
}
