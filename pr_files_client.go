package main

import (
	"fmt"
	"net/http"
)

type PrFilesClient struct {
	request *http.Request
}

func NewPrFilesClient(apiUrl string, token string, owner string, repo string, pullRequestNumber string) *PrFilesClient {
	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%s/files", apiUrl, owner, repo, pullRequestNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	return &PrFilesClient{
		request: req,
	}
}

func (client PrFilesClient) getPrFiles() (*http.Response, error) {
	resp, err := http.DefaultClient.Do(client.request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
