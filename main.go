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
	fmt.Println("token"+token, "repo"+repo, "pullNumber"+pullNumber)
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

	fmt.Println(string(body))
	var files DiffEntries
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, err
	}

	return files, nil
}

func comment(message string) {
	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY")
	pullNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")
	reviewer := NewPrReviewer("https://api.github.com", token, repo, pullNumber)
	err := reviewer.comment(message)
	if err != nil {
		fmt.Println("Error commenting on PR:", err)
		return
	}
}
func main() {
	files, err := getPRFiles()
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	comment("Hello from Go!")
}
