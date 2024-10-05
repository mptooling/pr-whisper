package domain

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

type CommentCondition func(DiffEntry, DiffEntries) bool

type trigger struct {
	checks []CommentCondition
}

type GenericWhisperer struct {
	Name     string
	Trigger  trigger
	Severity int
	Message  string
}

type Comment struct {
	WhisperName string
	Content     string
	Severity    int
	FilePath    string
	Position    int
	CommitID    string
}

type PRReview struct {
	Body  string `json:"body"`
	Event string `json:"event"`
}

type WhisperConfig struct {
	Whispers []WhisperConfigItem `yaml:"whispers"`
}

type WhisperConfigItem struct {
	Name     string    `yaml:"name"`
	Triggers []Trigger `yaml:"triggers"`
	Severity string    `yaml:"severity"`
	Message  string    `yaml:"message"`
}

type Trigger struct {
	Check    string `yaml:"check"`
	Contains string `yaml:"contains,omitempty"`
}
