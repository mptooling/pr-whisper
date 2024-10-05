package business

import (
	"github.com/mptooling/pr-whisper/internal/domain"
	"strings"
)

type GenericWhispererFactory struct{}

func NewGenericWhispererFactory() *GenericWhispererFactory {
	return &GenericWhispererFactory{}
}

func (w GenericWhispererFactory) MakeGenericWhispers(config *domain.WhisperConfig) []*GenericWhisperer {
	whispers := make([]*GenericWhisperer, 0)
	for _, whisper := range config.Whispers {
		if whisper.Triggers == nil || len(whisper.Triggers) == 0 {
			continue
		}

		checks := make([]CommentCondition, 0)
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

		whispers = append(whispers, &GenericWhisperer{
			Name: whisper.Name,
			Trigger: trigger{
				checks: checks,
			},
			Severity: severity,
			Message:  whisper.Message,
		})
	}

	return whispers
}
