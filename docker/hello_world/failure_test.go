package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestFailureExample(t *testing.T) {
	tag := "gruntwork/docker-hello-world-example"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "./", buildOptions)

	opts := &docker.RunOptions{Command: []string{"cat", "/test2.txt"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "Hello, World!", output)
}
