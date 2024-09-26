package main

const (
	Note = iota
	Tip
	Important
	Warning
	Caution
)

type DiffEntry struct {
	Sha              string `json:"sha"`
	Filename         string `json:"filename"`
	Status           string `json:"status"`
	Additions        int    `json:"additions"`
	Deletions        int    `json:"deletions"`
	Changes          int    `json:"changes"`
	BlobURL          string `json:"blob_url"`
	RawURL           string `json:"raw_url"`
	ContentsURL      string `json:"contents_url"`
	Patch            string `json:"patch,omitempty"`
	PreviousFilename string `json:"previous_filename,omitempty"`
}

type DiffEntries []DiffEntry

type Comment struct {
	WhisperName string
	Content     string
	Type        int
	FilePath    string
	Position    int
	CommitID    string
}

type PRInspiration struct {
	Body  string `json:"body"`
	Event string `json:"event"`
}

type PRReview struct {
	Body     string            `json:"body"`
	Event    string            `json:"event"`
	Comments []PrReviewComment `json:"comments"`
}

type PrReviewComment struct {
	Path     string `json:"path"`
	Position int    `json:"position"`
	Body     string `json:"body"`
}
