package main

import "strings"

type CommentCondition func(DiffEntry, DiffEntries) bool

type trigger struct {
	checks       []CommentCondition
	fileStatuses []string
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
		checks := make([]CommentCondition, 0)
		for _, trigger := range whisper.Triggers {
			if trigger.Check == "filepath" {
				checks = append(checks, func(change DiffEntry, changes DiffEntries) bool {
					return strings.Contains(change.Filename, trigger.Contains)
				})
			}

			if trigger.Check == "no_filepath_in_pr_files" {
				checks = append(checks, func(change DiffEntry, changes DiffEntries) bool {
					for _, c := range changes {
						if strings.Contains(c.Filename, trigger.Contains) {
							return false
						}
					}
					return true
				})
			}
		}

		fileStatuses := make([]string, 0)
		for _, trigger := range whisper.Triggers {
			if trigger.FileStatuses != nil {
				fileStatuses = trigger.FileStatuses
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
				checks:       checks,
				fileStatuses: fileStatuses,
			},
			Severity: severity,
			Message:  whisper.Message,
		})
	}

	return whispers
}
