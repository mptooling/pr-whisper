package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func getPRFiles() ([]map[string]interface{}, error) {
	token := os.Getenv("GITHUB_TOKEN")
	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPO")
	pullNumber := os.Getenv("GITHUB_PULL_NUMBER")
	client := NewPrFilesClient("https://api.github.com", token, owner, repo, pullNumber)

	resp, err := client.getPrFiles()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var files []map[string]interface{}
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, err
	}

	return files, nil
}

func comment(message string) {
	reviewer := NewPrReviewer("https://api.github.com", os.Getenv("GITHUB_TOKEN"), os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), os.Getenv("GITHUB_PULL_NUMBER"))
	err := reviewer.comment(message)
	if err != nil {
		fmt.Println("Error commenting on PR:", err)
		return
	}
}
func main() {
	files, err := getPRFiles()
	if err != nil {
		fmt.Println("Error fetching PR files:", err)
		return
	}
	for _, file := range files {
		fmt.Println(file["filename"])
	}

	comment("Hello from Go!")
}
