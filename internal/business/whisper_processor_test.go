package business

import (
	"errors"
	"github.com/mptooling/pr-whisper/internal/domain"
	"testing"
)

type MockReviewer struct {
	err error
}

func (mr *MockReviewer) Comment(comments []*domain.Comment) error {
	return nil
}

func (mr *MockReviewer) comment(comments []*domain.Comment) error {
	return mr.err
}

func ValidCheck(change domain.DiffEntry, changes domain.DiffEntries) bool {
	return true
}

func InvalidCheck(change domain.DiffEntry, changes domain.DiffEntries) bool {
	return false
}

func ProcessWhispers_ValidChanges(t *testing.T) {
	whisper := &domain.GenericWhisperer{
		Name:     "Valid Whisper",
		Message:  "Valid message",
		Severity: domain.Caution,
		Trigger: domain.WhisperTrigger{
			Checks: []domain.CommentCondition{ValidCheck},
		},
	}
	reviewer := &MockReviewer{}
	processor := NewWhisperProcessor([]*domain.GenericWhisperer{whisper}, reviewer)

	changes := domain.DiffEntries{
		{Filename: "file1.go", Sha: "abc123"},
	}

	processor.ProcessWhispers(changes)
}

func ProcessWhispers_InvalidChanges(t *testing.T) {
	whisper := &domain.GenericWhisperer{
		Name:     "Invalid Whisper",
		Message:  "Invalid message",
		Severity: domain.Caution,
		Trigger: domain.WhisperTrigger{
			Checks: []domain.CommentCondition{ValidCheck},
		},
	}
	reviewer := &MockReviewer{}
	processor := NewWhisperProcessor([]*domain.GenericWhisperer{whisper}, reviewer)

	changes := domain.DiffEntries{
		{Filename: "file1.go", Sha: "abc123"},
	}

	processor.ProcessWhispers(changes)
}

func ProcessWhispers_ReviewerError(t *testing.T) {
	whisper := &domain.GenericWhisperer{
		Name:     "Valid Whisper",
		Message:  "Valid message",
		Severity: domain.Caution,
		Trigger: domain.WhisperTrigger{
			Checks: []domain.CommentCondition{ValidCheck},
		},
	}
	reviewer := &MockReviewer{err: errors.New("reviewer error")}
	processor := NewWhisperProcessor([]*domain.GenericWhisperer{whisper}, reviewer)

	changes := domain.DiffEntries{
		{Filename: "file1.go", Sha: "abc123"},
	}

	processor.ProcessWhispers(changes)
}
