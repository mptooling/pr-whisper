package main

import "fmt"

type WhisperProcessor struct {
	whisperPool *WhisperPool
	reviewer    *PrReviewer
}

func NewWhisperProcessor(whisperPool *WhisperPool, reviewer *PrReviewer) *WhisperProcessor {
	return &WhisperProcessor{
		whisperPool: whisperPool,
		reviewer:    reviewer,
	}
}

func (wp *WhisperProcessor) ProcessWhispers(changes DiffEntries) {
	comments := make([]*Comment, 0)
	for _, change := range changes {
		for _, whisper := range wp.whisperPool.GetWhispers() {
			comment := whisper.Process(change, changes)
			if comment != nil && comment.Content != "" {
				comments = append(comments, comment)
			}
		}

	}

	err := wp.reviewer.comment(comments)
	if err != nil {
		fmt.Println("Error commenting on PR:", err)

		return
	}
}
