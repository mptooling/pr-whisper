package main

import (
	"fmt"
	"strings"
)

type LaravelResourceChangedWhisper struct{}

func NewLaravelResourceChangedWhisper() *LaravelResourceChangedWhisper {
	return &LaravelResourceChangedWhisper{}
}

func (w *LaravelResourceChangedWhisper) Process(change DiffEntry) string {
	if false == strings.Contains(change.Filename, "app/Http/Resources") {
		fmt.Println("Skipping whisper for file:", change.Filename)

		return ""
	}

	return "Please make sure to update the [specification.yaml](https://github.com/Purpose-Green/platform-api/blob/main/resources/apidoc/specification.yaml) file."
}
