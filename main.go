package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/docker/buildx/build"
	"github.com/docker/buildx/builder"
	_ "github.com/docker/buildx/driver/remote"
	"github.com/docker/buildx/util/dockerutil"
	xprogress "github.com/docker/buildx/util/progress"
)

var (
	builderx       string
	imgRef         string
	dockerfilePath string
)

func init() {
	flag.StringVar(&builderx, "builder", "default", "buildx builder")
	flag.StringVar(&imgRef, "image", "", "Image reference")
	flag.StringVar(&dockerfilePath, "dockerfile", "./Dockerfile", "Path to Dockerfile")
}

func main() {
	flag.Parse()

	if len(imgRef) == 0 {
		fmt.Fprintln(os.Stderr, "missing image reference")
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "missing context path")
		os.Exit(1)
	}
	ctxPath := args[0]

	dockerCli, err := newDockerCLI()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create docker cli: %v", err)
		os.Exit(1)
	}

	buildOptions, err := NewBuildOptions(imgRef, ctxPath, dockerfilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "build options failed: %v", err)
		os.Exit(1)
	}

	opts := map[string]build.Options{
		imgRef: buildOptions,
	}

	b, err := builder.New(dockerCli, builder.WithName(builderx))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create new builder: %v", err)
		os.Exit(1)
	}

	nodes, err := b.LoadNodes(context.TODO(), false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed loading nodes: %v", err)
		os.Exit(1)
	}

	if len(nodes) == 0 {
		fmt.Fprintln(os.Stderr, "no builder nodes found")
		os.Exit(1)
	}

	progressCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w, err := xprogress.NewPrinter(progressCtx, dockerCli.Out(), os.Stdout, "auto")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed creating printer: %v", err)
		os.Exit(1)
	}

	_, err = build.Build(context.TODO(), nodes, opts, dockerutil.NewClient(dockerCli), dockerCli.ConfigFile().Filename, w)
	if err != nil {
		fmt.Fprintf(os.Stderr, "build failed: %v", err)
		os.Exit(1)
	}
}
