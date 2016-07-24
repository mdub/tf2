package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/docker/distribution/context"
	"github.com/docker/distribution/reference"
	"github.com/docker/distribution/registry/client"
	"github.com/docker/distribution/registry/client/auth"
	ct "github.com/docker/distribution/registry/client/transport"
)

func main() {

	// hard-coded POC - list tags for "debian"

	repositoryName := "debian"

	repositoryRef, err := reference.ParseNamed(repositoryName)
	if err != nil {
		abort(2, err)
	}

	fmt.Println(repositoryRef)

	schemedHost := "https://registry-1.docker.io"

	transport, err := makeTransport(repositoryName)
	if err != nil {
		abort(3, err)
	}

	ctx := context.Background()

	repoClient, err := client.NewRepository(ctx, repositoryRef, schemedHost, transport)
	if err != nil {
		abort(3, err)
	}

	tags, err := repoClient.Tags(ctx).All(ctx)
	if err != nil {
		abort(4, err)
	}

	fmt.Println(tags)

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

func makeTransport(repositoryName string) (http.RoundTripper, error) {
	transport := http.DefaultTransport
	transport = NewDebugTransport(transport, os.Stderr)

	challengeManager := auth.NewSimpleChallengeManager()
	credentialStore := dumbCredentialStore{"", ""}
	tokenHandler := auth.NewTokenHandler(transport, credentialStore, repositoryName, "pull")
	basicHandler := auth.NewBasicHandler(credentialStore)
	authorizer := auth.NewAuthorizer(challengeManager, tokenHandler, basicHandler)

	transport = ct.NewTransport(transport, authorizer)

	return transport, nil
}

func abort(status int, message interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s", message)
	os.Exit(status)
}

type dumbCredentialStore struct {
	username string
	password string
}

func (dcs dumbCredentialStore) Basic(*url.URL) (string, string) {
	return dcs.username, dcs.password
}

func (dcs dumbCredentialStore) RefreshToken(*url.URL, string) string {
	return ""
}

func (dcs dumbCredentialStore) SetRefreshToken(realm *url.URL, service, token string) {
}
