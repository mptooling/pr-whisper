package adapters

import (
	"github.com/mptooling/pr-whisper/internal/domain"
	"net/http"
)

type PrReviewer interface {
	Comment([]*domain.Comment) error
}

type PrFilesClient interface {
	GetPrFiles() (*http.Response, error)
}
