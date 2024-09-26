package main

type WhisperPool struct {
	Whispers []Whisper
}

func NewWhisperPool() *WhisperPool {
	return &WhisperPool{}
}

func (wp *WhisperPool) AddWhisper(w Whisper) {
	wp.Whispers = append(wp.Whispers, w)
}

func (wp *WhisperPool) GetWhispers() []Whisper {
	return wp.Whispers
}
