// list-repos utilizes the go-github library to list all the repositories
// that the authorized user has access to.
//
// The following environment variables must be set:
//
// GITHUB_BASE_URL - base URL to the target enterprise github instance, 
// example: https://github.my-company.com/
//
// GITHUB_AUTH_TOKEN - authentication token to the target github instance

package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Environment variable 'GITHUB_AUTH_TOKEN' must be set.")
	}
	
	githubBaseUrl  := os.Getenv("GITHUB_BASE_URL")
	if githubBaseUrl == "" {
		log.Fatal("Environment variable 'GITHUB_BASE_URL' must be set.")
	}

	ctx := context.Background()

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient := oauth2.NewClient(ctx, tokenSource)

	client, err := github.NewEnterpriseClient(githubBaseUrl, githubBaseUrl, httpClient)
	if err != nil {
		log.Fatal(err)
	}

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	for {
		repos, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			log.Fatal(err)
		}

		for i, repo := range repos {
			fmt.Printf("%v.%v. %v\n", opt.Page, i+1, repo.GetName())
		}

		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}
}
