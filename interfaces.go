package main

type Whisper interface {
	Process(change DiffEntry) string
}
