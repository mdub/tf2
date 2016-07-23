package main

import (
	"fmt"
	"os"

	"github.com/docker/docker/reference"
	"github.com/docker/docker/registry"
)

func main() {

	// hard-coded POC - list tags for "debian"

	repositoryName := "debian"

	registryService := registry.NewService(registry.ServiceOptions{})

	repositoryRef, err := reference.ParseNamed(repositoryName)
	if err != nil {
		abort(2, err)
	}

	repositoryInfo, err := registryService.ResolveRepository(repositoryRef)
	if err != nil {
		abort(2, err)
	}

	fmt.Println(repositoryInfo.Named)

	// repository, confirmedV2, err = distribution.NewV2Repository(context.Foreground(), repoInfo, registryServvice.A, http.Header{}, authConfig, "pull")

}

func abort(status int, message interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s", message)
	os.Exit(status)
}
