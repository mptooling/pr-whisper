package main

import (
	"github.com/mptooling/pr-whisper/internal/business"
	cfg "github.com/mptooling/pr-whisper/internal/config"
	"os"
)

func main() {
	token := os.Getenv("GH_AUTH_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY")
	pullNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")

	config := cfg.NewConfig("whispers.yaml")
	whispersConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	whisperer := business.NewPrWhisper(token, repo, pullNumber, whispersConfig)
	if err := whisperer.Whisper(); err != nil {
		panic(err)
	}
}
