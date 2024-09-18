package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type PrReviewer struct {
	url     string
	headers map[string]string
}

func NewPrReviewer(apiUrl, token, repo, pullRequestNumber string) *PrReviewer {
	url := fmt.Sprintf("%s/repos/%s/pulls/%s/reviews", apiUrl, repo, pullRequestNumber)
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"Authorization":        "Bearer " + token,
		"X-GitHub-Api-Version": "2022-11-28",
	}

	return &PrReviewer{
		url:     url,
		headers: headers,
	}
}

func (client PrReviewer) comment(message string) error {
	jsonData := `{"body":"> [!IMPORTANT]`
	for _, line := range strings.Split(message, "\n") {
		if line == "" {
			continue
		}
		jsonData += `\n> ` + line
	}
	jsonData += `","event": "COMMENT"}`

	req, err := http.NewRequest("POST", client.url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return err
	}

	for key, value := range client.headers {
		req.Header.Set(key, value)
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

	if resp.Status != "200 OK" {
		return fmt.Errorf("error commenting on PR: %s", resp.Status)
	}

	return nil
}
