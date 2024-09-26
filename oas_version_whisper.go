package main

import (
	"fmt"
	"strings"
)

type OasVersionWhisper struct {
	name string
}

func NewOasVersionWhisper() *OasVersionWhisper {
	return &OasVersionWhisper{
		name: "OAS Version Updated Check",
	}
}

func (w *OasVersionWhisper) Process(change DiffEntry, changes DiffEntries) *Comment {
	fmt.Println("Processing OasVersionWhisper")
	if false == strings.Contains(change.Filename, "specification.yaml") {
		return nil
	}

	if change.Status != "modified" {
		fmt.Println("OasVersionWhisper: nothing to do")
		return nil
	}

	for _, change := range changes {
		if false != strings.Contains(change.Patch, "openapi:") {
			fmt.Println("OasVersionWhisper: attaboy! You have changed the version!")
			return nil
		}
	}

	content := "🤫 Psst... Looks like **specification.yaml** file has been modified."
	content += "\n\n- [ ] Update **openapi** version in the resources/apidoc/specification.yaml file."

	comment := &Comment{
		WhisperName: w.name,
		Content:     content,
		Type:        Important,
		FilePath:    change.Filename,
		Position:    1,
		CommitID:    change.Sha,
	}

	return comment
}
