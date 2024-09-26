package main

import (
	"fmt"
	"strings"
)

type ApiBcBreakWhisper struct {
	name string
}

func NewApiBcBreakWhisper() *ApiBcBreakWhisper {
	return &ApiBcBreakWhisper{
		name: "BC Break Check",
	}
}

func (w *ApiBcBreakWhisper) Process(change DiffEntry, changes DiffEntries) *Comment {
	fmt.Println("Processing ApiBcBreakWhisper")
	comment := w.processResources(change)
	if comment != nil {
		return comment
	}

	comment = w.processControllers(change)
	if comment != nil {
		return comment
	}

	comment = w.processRequest(change)
	if comment != nil {
		return comment
	}

	comment = w.processRouting(change)
	if comment != nil {
		return comment
	}

	return nil
}

func (w *ApiBcBreakWhisper) processResources(change DiffEntry) *Comment {
	if false == strings.Contains(change.Filename, "app/Http/Resources") {
		return nil
	}

	if change.Status != "modified" && change.Status != "removed" {
		fmt.Println("ApiBcBreakWhisper: nothing to do")
		return nil
	}

	var commentType int
	var content string
	switch change.Status {
	case "modified":
		commentType = Warning
		content = "⚠️ Psst... Looks like this Resource file has been modified."
	case "removed":
		commentType = Caution
		content = "❌ This Resource has been removed."
	}

	content += "\n\n"
	content += "- [ ] Make sure there are no breaking changes.\n"
	content += "- [ ] Align this change with the FE team before merge.\n"

	return &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        commentType,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}

func (w *ApiBcBreakWhisper) processControllers(change DiffEntry) *Comment {
	if false == strings.Contains(change.Filename, "app/Http/Controllers") {
		return nil
	}

	if change.Status != "modified" && change.Status != "removed" {
		return nil
	}

	var commentType int
	var content string
	switch change.Status {
	case "modified":
		commentType = Warning
		content = "⚠️ Psst... Looks like this Controller file has been modified"
	case "removed":
		commentType = Caution
		content = "❌ This Controller has been removed."
	}

	content += "\n\n"
	content += "- [ ] Make sure there are no breaking changes.\n"
	content += "- [ ] Align this change with the FE team before merge.\n"

	return &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        commentType,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}

func (w *ApiBcBreakWhisper) processRequest(change DiffEntry) *Comment {
	if false == strings.Contains(change.Filename, "app/DataTransferObjects") {
		return nil
	}

	if false == strings.Contains(change.Filename, "Request") && false == strings.Contains(change.Filename, "Data") {
		return nil
	}

	if change.Status != "modified" && change.Status != "removed" {
		return nil
	}

	var commentType int
	var content string
	switch change.Status {
	case "modified":
		commentType = Warning
		content = "⚠️ Psst... Looks like this Request file has been modified."
	case "removed":
		commentType = Caution
		content = "❌ This Request has been removed."
	}

	content += "\n\n"
	content += "- [ ] Make sure there are no breaking changes.\n"
	content += "- [ ] Align this change with the FE team before merge.\n"

	return &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        commentType,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}

func (w *ApiBcBreakWhisper) processRouting(change DiffEntry) *Comment {
	if change.Filename != "routes/api.php" {
		return nil
	}

	if change.Status != "modified" {
		return nil
	}

	content := "⚠️ Looks like our routing file has been changed. This might be a breaking change. Don't be shy, talk to the FE team before merging this beauty 😉"
	content += "\n\n"
	content += "- [ ] Make sure there are no breaking changes.\n"
	content += "- [ ] Align this change with the FE team before merge.\n"

	return &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        Warning,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}
