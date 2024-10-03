package main

import "fmt"

type WhisperProcessor struct {
	whisperPool []*GenericWhisperer
	reviewer    *PrReviewer
}

func NewWhisperProcessor(whisperPool []*GenericWhisperer, reviewer *PrReviewer) *WhisperProcessor {
	return &WhisperProcessor{
		whisperPool: whisperPool,
		reviewer:    reviewer,
	}
}

func (wp *WhisperProcessor) ProcessWhispers(changes DiffEntries) {
	comments := make([]*Comment, 0)
	for _, change := range changes {
		wp.processChange(change, changes, comments)
	}

	err := wp.reviewer.comment(comments)
	if err != nil {
		fmt.Println("Error commenting on PR:", err)

		return
	}
}

func (wp *WhisperProcessor) processChange(change DiffEntry, changes DiffEntries, comments []*Comment) {
	for _, whisper := range wp.whisperPool {
		wp.runWhisperer(whisper, change, changes, comments)
	}
}

func (wp *WhisperProcessor) runWhisperer(w *GenericWhisperer, change DiffEntry, changes DiffEntries, comments []*Comment) {
	for _, check := range w.Trigger.checks {
		if false == check(change, changes) {
			return
		}
	}

	comments = append(comments, &Comment{
		WhisperName: w.Name,
		Content:     w.Message,
		Type:        w.Severity,
		FilePath:    change.Filename,
		Position:    0,
		CommitID:    change.Sha,
	})
}
