package business

import (
	"github.com/mptooling/pr-whisper/internal/adapters"
	"github.com/mptooling/pr-whisper/internal/domain"
)

type whisperProcessor struct {
	whisperPool []*domain.GenericWhisperer
	reviewer    adapters.PrReviewer
}

func NewWhisperProcessor(whisperPool []*domain.GenericWhisperer, reviewer adapters.PrReviewer) WhisperProcessor {
	return &whisperProcessor{
		whisperPool: whisperPool,
		reviewer:    reviewer,
	}
}

func (wp *whisperProcessor) ProcessWhispers(changes domain.DiffEntries) error {
	comments := make([]*domain.Comment, 0)
	for _, change := range changes {
		for _, c := range wp.processChange(change, changes) {
			comments = append(comments, c)
		}
	}

	err := wp.reviewer.Comment(comments)
	if err != nil {
		return err
	}

	return nil
}

func (wp *whisperProcessor) processChange(change domain.DiffEntry, changes domain.DiffEntries) []*domain.Comment {
	comments := make([]*domain.Comment, 0)
	for _, whisper := range wp.whisperPool {
		c := wp.runWhisperer(whisper, change, changes)
		if c != nil {
			comments = append(comments, wp.runWhisperer(whisper, change, changes))
		}
	}

	return comments
}

func (wp *whisperProcessor) runWhisperer(w *domain.GenericWhisperer, change domain.DiffEntry, changes domain.DiffEntries) *domain.Comment {
	if len(w.Trigger.Checks) == 0 {
		return nil
	}

	for _, check := range w.Trigger.Checks {
		if false == check(change, changes) {
			return nil
		}
	}

	return &domain.Comment{
		WhisperName: w.Name,
		Content:     w.Message,
		Severity:    w.Severity,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}
