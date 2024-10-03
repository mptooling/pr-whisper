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

	config := NewConfig("whispers.yaml")
	whispersConfig, err := config.loadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	factory := NewGenericWhispererFactory()
	whispers := factory.MakeGenericWhispers(whispersConfig)

	reviewer := NewPrReviewer("https://api.github.com", token, repo, pullNumber)

	files, err := getPRFiles(token, repo, pullNumber)
	if err != nil {
		fmt.Println("Error getting PR files:", err)
		return
	}

	processor := NewWhisperProcessor(whispers, reviewer)
	processor.ProcessWhispers(files)
}
