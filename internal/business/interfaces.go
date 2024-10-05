package business

import "github.com/mptooling/pr-whisper/internal/domain"

type PrWhisper interface {
	Whisper() error
}

type WhisperProcessor interface {
	ProcessWhispers(changes domain.DiffEntries) error
}
