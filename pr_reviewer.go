package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type PrReviewer struct {
	url     string
	headers map[string]string
}

func NewPrReviewer(apiUrl string, token string, repo string, pullRequestNumber string) *PrReviewer {
	url := fmt.Sprintf("%s/repos/%s/pulls/%s/reviews", apiUrl, repo, pullRequestNumber)
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"Authorization":        "Bearer " + token,
		"X-GitHub-Api-Version": "2022-11-28",
		"Content-Type":         "application/json",
	}

	return &PrReviewer{
		url:     url,
		headers: headers,
	}
}

func (client PrReviewer) comment(message string) error {
	jsonData := `{"body":` + message + `}`

	req, err := http.NewRequest("POST", client.url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
