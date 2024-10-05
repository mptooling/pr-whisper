package main

import (
	"encoding/json"
	"fmt"
	"github.com/mptooling/pr-whisper/internal/adapters"
	"github.com/mptooling/pr-whisper/internal/business"
	config2 "github.com/mptooling/pr-whisper/internal/config"
	"github.com/mptooling/pr-whisper/internal/domain"
	"io"
	"os"
)

// todo :: move to client adapter
func getPRFiles(token, repo, pullRequestNumber string) (domain.DiffEntries, error) {
	client := adapters.NewPrFilesClient("https://api.github.com", token, repo, pullRequestNumber)

	resp, err := client.GetPrFiles()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var files domain.DiffEntries
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	token := os.Getenv("GH_AUTH_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY")
	pullNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")

	config := config2.NewConfig("whispers.yaml")
	whispersConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	factory := business.NewGenericWhispererFactory()

	whispers := factory.MakeGenericWhispers(whispersConfig)

	reviewer := adapters.NewPrReviewer("https://api.github.com", token, repo, pullNumber)

	processor := business.NewWhisperProcessor(whispers, reviewer)

	files, err := getPRFiles(token, repo, pullNumber)
	if err != nil {
		fmt.Println("Error getting PR files:", err)
		return
	}

	processor.ProcessWhispers(files)
}
