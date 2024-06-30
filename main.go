package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/specgen"
)

func main() {
	// Get Podman socket location
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		runtimeDir = "/var/run"
	}

	socketPath := runtimeDir + "/docker.sock"

	// Connect to Podman socket
	connText, err := bindings.NewConnection(context.Background(), "unix://"+socketPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to Podman: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to Podman")

	// Pull Busybox image (Sample 1)
	fmt.Println("Pulling Alpine Nginx image...")
	rawImage := "quay.io/libpod/alpine_nginx"
	_, err = images.Pull(connText, rawImage, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// List Images
	imageSummary, err := images.List(connText, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var names []string
	for _, image := range imageSummary {
		names = append(names, image.Names...)
	}

	fmt.Println("Images:")
	fmt.Println(names)

	// Create and start a container
	s := specgen.NewSpecGenerator(rawImage, false)
	s.Name = "foobar"
	createResponse, err := containers.CreateWithSpec(connText, s, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Container created with ID:", createResponse.ID)

	fmt.Println("Starting container...")
	if err := containers.Start(connText, createResponse.ID, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Container started")

	// Container list
	containerLatestList, err := containers.List(connText, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Containers: %s\n", fmt.Sprint(containerLatestList))

	// Container inspect
	ctrData, err := containers.Inspect(connText, createResponse.ID, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Container Inspect Name: %s\n", ctrData.ImageName)
	fmt.Printf("Container Inspect Status: %s\n", ctrData.State.Status)

	// Container stop
	fmt.Println("Stopping container...")
	if err := containers.Stop(connText, createResponse.ID, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctrData, err = containers.Inspect(connText, createResponse.ID, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Container Inspect Status: %s\n", ctrData.State.Status)
}
