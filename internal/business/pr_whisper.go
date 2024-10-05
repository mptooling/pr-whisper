package business

import (
	"encoding/json"
	"fmt"
	"github.com/mptooling/pr-whisper/internal/adapters"
	"github.com/mptooling/pr-whisper/internal/domain"
	"io"
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
	factory := NewGenericWhispererFactory() // todo :: move to factory

	whispers := factory.MakeGenericWhispers(p.config)

	reviewer := adapters.NewPrReviewer(ghApiUrl, p.token, p.repository, p.prNumber)

	processor := NewWhisperProcessor(whispers, reviewer) // processor might be not needed in the way it currently is

	files, err := p.getPRFiles()
	if err != nil {
		fmt.Println("Error getting PR files:", err)

		return err
	}

	processor.ProcessWhispers(files) // todo :: handle error

	return nil
}

// todo :: move to client adapter
func (p prWhisper) getPRFiles() (domain.DiffEntries, error) {
	client := adapters.NewPrFilesClient("https://api.github.com", p.token, p.repository, p.prNumber)

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
