package main

import (
	"path/filepath"

	"github.com/docker/cli/cli/command"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/docker/docker/builder/remotecontext/urlutil"
)

func dockerFilePath(ctxName string, dockerfile string) string {
	if urlutil.IsGitURL(ctxName) || filepath.IsAbs(dockerfile) {
		return dockerfile
	}
	return filepath.Join(ctxName, dockerfile)
}

func newDockerCLI() (*command.DockerCli, error) {
	dockerCli, err := command.NewDockerCli()
	if err != nil {
		return nil, err
	}
	if err = dockerCli.Initialize(cliflags.NewClientOptions()); err != nil {
		return nil, err
	}
	return dockerCli, nil
}
