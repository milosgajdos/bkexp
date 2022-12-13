package main

import (
	"github.com/docker/buildx/build"
	"github.com/moby/buildkit/client"
)

func NewBuildOptions(image, ctxPath, dockerfilePath string) (build.Options, error) {
	cacheFrom := []client.CacheOptionsEntry{
		{Type: "registry", Attrs: map[string]string{"ref": image}},
	}

	exports := []client.ExportEntry{
		{Type: "docker", Attrs: map[string]string{"load": "true"}},
	}

	return build.Options{
		Inputs: build.Inputs{
			ContextPath:    ctxPath,
			DockerfilePath: dockerFilePath(ctxPath, dockerfilePath),
		},
		CacheFrom: cacheFrom,
		Exports:   exports,
		Tags:      []string{image},
	}, nil
}
