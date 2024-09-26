package main

type Whisper interface {
	Process(change DiffEntry, changes DiffEntries) *Comment
}
