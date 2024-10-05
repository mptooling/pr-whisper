package main

import (
	"fmt"
	"github.com/mptooling/pr-whisper/internal/business"
	config2 "github.com/mptooling/pr-whisper/internal/config"
	"os"
)

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

	whisperer := business.NewPrWhisper(token, repo, pullNumber, whispersConfig)
	if err := whisperer.Whisper(); err != nil {
		panic(err)
	}
}
