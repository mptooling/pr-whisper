package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func getPRFiles() (DiffEntries, error) {
	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY")
	pullNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")
	client := NewPrFilesClient("https://api.github.com", token, repo, pullNumber)

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

func comment(message string) error {
	token := os.Getenv("GH_AUTH_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY")
	pullNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")
	reviewer := NewPrReviewer("https://api.github.com", token, repo, pullNumber)
	err := reviewer.comment(message)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	files, err := getPRFiles()
	if err != nil {
		fmt.Println("Error getting PR files:", err)
	}

	for _, file := range files {
		fmt.Printf("File: %s. Status: %s \n", file.Filename, file.Status)
	}

	err = comment("Hello from Go!")
	if err != nil {
		fmt.Println("Error commenting on PR:", err)
	}
}
