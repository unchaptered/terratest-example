package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
)

func TestSuccessExample(t *testing.T) {
	tag := "sample"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}
	docker.Build(t, "./", buildOptions)

	// docker run --detach -P sample
	runOpts := &docker.RunOptions{
		Detach: true,
		OtherOptions: []string{"-P"},
	}
	containerID := docker.Run(t, tag, runOpts)
	
	// docker stop $containerID
	stopOpts := &docker.StopOptions{}
	defer docker.Stop(t, []string{containerID}, stopOpts)


}   