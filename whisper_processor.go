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

func (wp *WhisperProcessor) ProcessWhispers(change DiffEntry) {
	for _, whisper := range wp.whisperPool.GetWhispers() {
		msg := whisper.Process(change)
		err := wp.reviewer.comment(msg)
		if err != nil {
			fmt.Println("Error commenting on PR:", err)
			return
		}
	}
}
