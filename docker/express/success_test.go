package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestSuccessExample(t *testing.T) {
	tag := "unchaptered/docker-express"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "./", buildOptions)

	runOptions := &docker.RunOptions{
		Detach: true,
	}
	output := docker.Run(t, tag, runOptions)
	assert.Equal(t, output, "hh")
}   