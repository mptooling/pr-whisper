package adapters

import (
	"github.com/mptooling/pr-whisper/internal/domain"
)

type PrReviewer interface {
	Comment([]*domain.Comment) error
}

type PrFilesClient interface {
	GetPrFiles() (domain.DiffEntries, error)
}
