package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	"github.com/docker/docker/cliconfig"
	"github.com/docker/docker/reference"
	"github.com/docker/docker/registry"
	registrytypes "github.com/docker/engine-api/types/registry"
)

func main() {

	// hard-coded POC - list tags for "debian"

	repositoryName := "debian"

	repositoryRef, err := reference.ParseNamed(repositoryName)
	if err != nil {
		abort(2, err)
	}

	fmt.Println(repositoryRef)

	cliConfig, err := cliconfig.Load(cliconfig.ConfigDir())
	if err != nil {
		abort(2, err)
	}

	indexInfo := &registrytypes.IndexInfo{
		Name: repositoryRef.Hostname(),
	}

	authConfig := registry.ResolveAuthConfig(cliConfig.AuthConfigs, indexInfo)
	fmt.Println(authConfig)

	registryService := registry.NewService(registry.ServiceOptions{})
	fmt.Println(registryService)

	status, token, err := registryService.Auth(context.TODO(), &authConfig, "tf2")

	fmt.Println(status, token)

	//
	// authConfig := registry.ResolveAuthConfig(cliConfig.AuthConfigs, repositoryInfo.Index)
	//
	// schemedHost := "https://registry-1.docker.io"
	//
	// transport, err := makeTransport(repositoryName)
	// if err != nil {
	// 	abort(3, err)
	// }
	//
	// ctx := context.Background()
	//
	// repoClient, err := client.NewRepository(ctx, repositoryRef, schemedHost, transport)
	// if err != nil {
	// 	abort(3, err)
	// }
	//
	// tags, err := repoClient.Tags(ctx).All(ctx)
	// if err != nil {
	// 	abort(4, err)
	// }
	//
	// fmt.Println(tags)

	// cliConfig, err := cliconfig.Load(cliconfig.ConfigDir())
	// if err != nil {
	// 	abort(2, err)
	// }
	//
	// authConfig := registry.ResolveAuthConfig(cliConfig.AuthConfigs, repositoryInfo.Index)
	//
	// metaHeaders := http.Header{}
	//

}

func abort(status int, message interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s", message)
	os.Exit(status)
}
