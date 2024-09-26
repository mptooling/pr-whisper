package main

import (
	"fmt"
	"strings"
)

type OasConsistencyWhisper struct {
	name string
}

func NewOasConsistencyWhisper() *OasConsistencyWhisper {
	return &OasConsistencyWhisper{
		name: "OAS Consistency Check",
	}
}

func (w *OasConsistencyWhisper) Process(change DiffEntry, changes DiffEntries) *Comment {
	comment := w.processResource(change, changes)
	if comment != nil {
		return comment
	}

	comment = w.processController(change, changes)
	if comment != nil {
		return comment
	}

	comment = w.processRequest(change, changes)
	if comment != nil {
		return comment
	}

	comment = w.processRouting(change, changes)
	if comment != nil {
		return comment
	}

	return nil
}

func (w *OasConsistencyWhisper) processResource(change DiffEntry, changes DiffEntries) *Comment {
	fmt.Println("Processing OasConsistencyWhisper")
	if false == strings.Contains(change.Filename, "app/Http/Resources") {
		return nil
	}

	if change.Status != "modified" && change.Status != "added" && change.Status != "removed" && change.Status != "changed" {
		fmt.Println("OasConsistencyWhisper: nothing to do")
		return nil
	}

	for _, change := range changes {
		if strings.Contains(change.Filename, "specification.yaml") {
			fmt.Println("OasConsistencyWhisper: specification already modified")

			return nil
		}
	}

	var commentType int
	var content string
	switch change.Status {
	case "modified":
		commentType = Important
		content = "🤫Psst... Looks like this Resource file has been modified. Just a heads-up — don't forget to update the **specification.yaml** file"
	case "added":
		commentType = Note
		content = "🌱 My oh my. Brand new Resource. Don't forget to update the **specification.yaml** file"
	case "removed":
		commentType = Caution
		content = "🪦 No Resource, no problem. Please, don't forget to update the **specification.yaml** file"
	}

	content += "\n\n - [ ] Update resources/apidoc/specification.yaml"

	comment := &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        commentType,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}

	return comment
}

func (w *OasConsistencyWhisper) processController(change DiffEntry, changes DiffEntries) *Comment {
	fmt.Println("Processing OasConsistencyWhisper")
	if false == strings.Contains(change.Filename, "app/Http/Controller") {
		return nil
	}

	if change.Status != "modified" && change.Status != "added" && change.Status != "removed" && change.Status != "changed" {
		return nil
	}

	for _, change := range changes {
		if strings.Contains(change.Filename, "specification.yaml") {
			return nil
		}
	}

	var commentType int
	var content string
	switch change.Status {
	case "modified":
		commentType = Important
		content = "🤫 Psst... Looks like this Controller file has been modified. Just a heads-up — don't forget to update the **specification.yaml** file"
	case "added":
		commentType = Note
		content = "🌱 Psst... Looks like this is a new Controller. Don't forget to update the **specification.yaml** file"
	case "removed":
		commentType = Caution
		content = "🪦 No Controller, no problem. Please, don't forget to update the **specification.yaml** file"
	}

	content += "\n\n - [ ] Update resources/apidoc/specification.yaml"

	comment := &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        commentType,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}

	return comment
}

func (w *OasConsistencyWhisper) processRequest(change DiffEntry, changes DiffEntries) *Comment {
	if false == strings.Contains(change.Filename, "app/DataTransferObjects") {
		return nil
	}

	if false == strings.Contains(change.Filename, "Request") {
		return nil
	}

	if false == strings.Contains(change.Filename, "Data") {
		return nil
	}

	if change.Status != "modified" && change.Status != "removed" {
		return nil
	}

	for _, change := range changes {
		if strings.Contains(change.Filename, "specification.yaml") {
			fmt.Println("OasConsistencyWhisper: specification already modified")

			return nil
		}
	}

	var commentType int
	var content string
	switch change.Status {
	case "modified":
		commentType = Important
		content = "🤫 Psst... Looks like this Request file has been modified. Just a heads-up — don't forget to update the **specification.yaml** file"
	case "added":
		commentType = Note
		content = "🌱 Psst... Looks like this is a new Request. Don't forget to update the **specification.yaml** file"
	case "removed":
		commentType = Caution
		content = "🪦 No Request, no problem. Please, don't forget to update the **specification.yaml** file"
	}

	content += "\n\n - [ ] Update resources/apidoc/specification.yaml"

	return &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        commentType,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}

func (w *OasConsistencyWhisper) processRouting(change DiffEntry, changes DiffEntries) *Comment {
	if false == strings.Contains(change.Filename, "routes/api.php") {
		return nil
	}

	if change.Status != "modified" {
		return nil
	}

	for _, change := range changes {
		if strings.Contains(change.Filename, "specification.yaml") {
			fmt.Println("OasConsistencyWhisper: specification already modified")

			return nil
		}
	}

	content := "🤫 Psst... Looks like this Request object has been modified. Please, don't forget to update the **specification.yaml** file"
	content += "\n\n - [ ] Update resources/apidoc/specification.yaml"

	return &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        Important,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}
}
