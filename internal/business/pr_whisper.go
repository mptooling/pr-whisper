package business

import (
	"github.com/mptooling/pr-whisper/internal/adapters"
	"github.com/mptooling/pr-whisper/internal/domain"
)

const ghApiUrl = "https://api.github.com"

type prWhisper struct {
	token      string
	repository string
	prNumber   string
	config     *domain.WhisperConfig
}

func NewPrWhisper(token, repository, prNumber string, config *domain.WhisperConfig) PrWhisper {
	return &prWhisper{
		token:      token,
		repository: repository,
		prNumber:   prNumber,
		config:     config,
	}
}

func (p prWhisper) Whisper() error {
	prDataClient := adapters.NewPrDataClient(ghApiUrl, p.token, p.repository, p.prNumber)
	files, err := prDataClient.GetPrFiles()
	if err != nil {
		return err
	}

	whispers := NewGenericWhispererBuilder().MakeGenericWhispers(p.config)
	reviewer := adapters.NewPrReviewer(ghApiUrl, p.token, p.repository, p.prNumber)

	return NewWhisperProcessor(whispers, reviewer).ProcessWhispers(files)
}
