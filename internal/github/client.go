package github

import (
	"context"
	"log"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

type Client struct {
	client *github.Client
}

func NewClient(appID, installationID int64) *Client {

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID authenticating with installation ID 99.
	itr, err := ghinstallation.NewKeyFromFile(tr, appID, installationID, "secret.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	return &Client{client: client}
}

func (c *Client) ListRepos() ([]*github.Repository, error) {
	var all []*github.Repository

	var nextPage int
	for next := true; next != false; {
		repos, rs, err := c.client.Apps.ListRepos(context.Background(), &github.ListOptions{Page: nextPage})
		if err != nil {
			return nil, err
		}

		all = append(all, repos...)

		next = rs.NextPage != 0
		nextPage = rs.NextPage
	}

	return all, nil
}
