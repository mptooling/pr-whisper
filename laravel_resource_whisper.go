package main

import (
	"fmt"
	"strings"
)

type LaravelResourceWhisper struct{}

func NewLaravelResourceChangedWhisper() *LaravelResourceWhisper {
	return &LaravelResourceWhisper{}
}

func (w *LaravelResourceWhisper) Process(change DiffEntry) string {
	if false == strings.Contains(change.Filename, "app/Http/Resources") {
		fmt.Println("Skipping whisper for file:", change.Filename)

		return ""
	}

	if change.Status == "modified" || change.Status == "added" || change.Status == "removed" || change.Status == "changed" {
		return "Please make sure to update the [specification.yaml](https://github.com/Purpose-Green/platform-api/blob/main/resources/apidoc/specification.yaml) file."
	}

	return ""
}
