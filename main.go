package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func getPRFiles(token, repo, pullRequestNumber string) (DiffEntries, error) {
	client := NewPrFilesClient("https://api.github.com", token, repo, pullRequestNumber)

	resp, err := client.getPrFiles()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var files DiffEntries
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	token := os.Getenv("GH_AUTH_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY")
	pullNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")

	files, err := getPRFiles(token, repo, pullNumber)
	if err != nil {
		fmt.Println("Error getting PR files:", err)
		return
	}

	wp := NewWhisperPool()
	wp.AddWhisper(NewOasConsistencyWhisper())
	wp.AddWhisper(NewApiBcBreakWhisper())
	wp.AddWhisper(NewOasVersionWhisper())
	processor := NewWhisperProcessor(wp, NewPrReviewer("https://api.github.com", token, repo, pullNumber))
	processor.ProcessWhispers(files)
}
