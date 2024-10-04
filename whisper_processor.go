package main

import "fmt"

type WhisperProcessor struct {
	whisperPool []*GenericWhisperer
	reviewer    PrClient
}

func NewWhisperProcessor(whisperPool []*GenericWhisperer, reviewer PrClient) *WhisperProcessor {
	return &WhisperProcessor{
		whisperPool: whisperPool,
		reviewer:    reviewer,
	}
}

func (wp *WhisperProcessor) ProcessWhispers(changes DiffEntries) {
	comments := make([]*Comment, 0)
	for _, change := range changes {
		for _, c := range wp.processChange(change, changes) {
			comments = append(comments, c)
		}
	}

	fmt.Println("collected comments:", comments)
	err := wp.reviewer.comment(comments)
	if err != nil {
		fmt.Println("Error commenting on PR:", err)

		return
	}
}

func (wp *WhisperProcessor) processChange(change DiffEntry, changes DiffEntries) []*Comment {
	comments := make([]*Comment, 0)
	for _, whisper := range wp.whisperPool {
		c := wp.runWhisperer(whisper, change, changes)
		if c != nil {
			comments = append(comments, wp.runWhisperer(whisper, change, changes))
		}
	}

	return comments
}

func (wp *WhisperProcessor) runWhisperer(w *GenericWhisperer, change DiffEntry, changes DiffEntries) *Comment {
	if len(w.Trigger.checks) == 0 {
		return nil
	}

	for _, check := range w.Trigger.checks {
		if false == check(change, changes) {
			return nil
		}
	}

	return &Comment{
		WhisperName: w.Name,
		Content:     w.Message,
		Severity:    w.Severity,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}
