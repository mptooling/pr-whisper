package business

import (
	"github.com/mptooling/pr-whisper/internal/domain"
	"strings"
)

type GenericWhispererBuilder struct{}

func NewGenericWhispererBuilder() *GenericWhispererBuilder {
	return &GenericWhispererBuilder{}
}

func (w GenericWhispererBuilder) MakeGenericWhispers(config *domain.WhisperConfig) []*domain.GenericWhisperer {
	whispers := make([]*domain.GenericWhisperer, 0)
	for _, whisper := range config.Whispers {
		if whisper.Triggers == nil || len(whisper.Triggers) == 0 {
			continue
		}

		checks := w.buildTriggers(whisper, make([]domain.CommentCondition, 0))
		severity := w.getSeverityLevel(whisper)

		whispers = append(whispers, &domain.GenericWhisperer{
			Name: whisper.Name,
			Trigger: domain.WhisperTrigger{
				Checks: checks,
			},
			Severity: severity,
			Message:  whisper.Message,
		})
	}

	return whispers
}

func (w GenericWhispererBuilder) getSeverityLevel(whisper domain.WhisperConfigItem) int {
	severity := domain.Note
	switch whisper.Severity {
	case "warning":
		severity = domain.Warning
	case "important":
		severity = domain.Important
	case "tip":
		severity = domain.Tip
	case "caution":
		severity = domain.Caution
	default:
		severity = domain.Note
	}
	return severity
}

func (w GenericWhispererBuilder) buildTriggers(whisper domain.WhisperConfigItem, checks []domain.CommentCondition) []domain.CommentCondition {
	for _, trigger := range whisper.Triggers {
		if trigger.Check == "filepath" {
			checks = append(checks, func(change domain.DiffEntry, changes domain.DiffEntries) bool {
				return strings.Contains(change.Filename, trigger.Contains)
			})

			continue
		}

		if trigger.Check == "file_not_in_pr" {
			filePath := trigger.Contains
			checks = append(checks, func(change domain.DiffEntry, changes domain.DiffEntries) bool {
				for _, c := range changes {
					if strings.Contains(c.Filename, filePath) {
						return false
					}
				}
				return true
			})

			continue
		}

		if trigger.Check == "file_status" {
			checks = append(checks, func(change domain.DiffEntry, changes domain.DiffEntries) bool {
				statuses := strings.Split(trigger.Contains, ",")
				for _, fileStatus := range statuses {
					if change.Status == fileStatus {
						return true
					}
				}
				return false
			})
		}
	}

	return checks
}
