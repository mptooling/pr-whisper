package main

import (
	"strings"
)

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

type GenericWhispererFactory struct{}

func NewGenericWhispererFactory() *GenericWhispererFactory {
	return &GenericWhispererFactory{}
}

func (w GenericWhispererFactory) MakeGenericWhispers(config *WhisperConfig) []*GenericWhisperer {
	whispers := make([]*GenericWhisperer, 0)
	for _, whisper := range config.Whispers {
		if whisper.Triggers == nil || len(whisper.Triggers) == 0 {
			continue
		}

		checks := make([]CommentCondition, 0)
		for _, trigger := range whisper.Triggers {
			if trigger.Check == "filepath" {
				checks = append(checks, func(change DiffEntry, changes DiffEntries) bool {
					return strings.Contains(change.Filename, trigger.Contains)
				})

				continue
			}

			if trigger.Check == "file_not_in_pr" {
				filePath := trigger.Contains
				checks = append(checks, func(change DiffEntry, changes DiffEntries) bool {
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
				checks = append(checks, func(change DiffEntry, changes DiffEntries) bool {
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

		severity := Note
		switch whisper.Severity {
		case "warning":
			severity = Warning
		case "important":
			severity = Important
		case "tip":
			severity = Tip
		case "caution":
			severity = Caution
		default:
			severity = Note
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
