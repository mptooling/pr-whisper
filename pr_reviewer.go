package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PrReviewer struct {
	url     string
	headers map[string]string
}

func NewPrReviewer(apiUrl, token, repo, pullRequestNumber string) *PrReviewer {
	url := fmt.Sprintf("%s/repos/%s/pulls/%s/reviews", apiUrl, repo, pullRequestNumber)
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"Authorization":        "Bearer " + token,
		"X-GitHub-Api-Version": "2022-11-28",
	}

	return &PrReviewer{
		url:     url,
		headers: headers,
	}
}

func (client PrReviewer) comment(comments []*Comment) error {
	if len(comments) == 0 {
		return nil
	}

	return client.commentWhispers(comments)
}

func (client PrReviewer) commentWhispers(comments []*Comment) error {
	var commentType string
	var cs []PrReviewComment
	commentTypeString := Note
	for _, c := range comments {
		if c.Type > commentTypeString {
			commentTypeString = c.Type
		}

		cs = append(cs, PrReviewComment{
			Path:     c.FilePath,
			Position: c.Position,
			Body:     c.Content,
		})
	}
	switch commentTypeString {
	case Important:
		commentType = "IMPORTANT"
	case Caution:
		commentType = "CAUTION"
	case Warning:
		commentType = "WARNING"
	case Tip:
		commentType = "TIP"
	default:
		commentType = "NOTE"
	}

	review := PRReview{
		Body:     `> [!` + commentType + `]` + "\n" + `> ` + client.randomIntroString(commentTypeString) + "\n",
		Event:    "COMMENT",
		Comments: cs,
	}

	jsonData, err := json.Marshal(review)
	if err != nil {
		return err
	}

	return client.send(jsonData)
}

func (client PrReviewer) send(jsonData []byte) error {
	req, err := http.NewRequest("POST", client.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	for key, value := range client.headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	if resp.Status != "200 OK" {
		return fmt.Errorf("error commenting on PR: %s", resp.Status)
	}

	return nil
}

func (client PrReviewer) randomIntroString(severityLevel int) string {
	switch severityLevel {
	case Caution:
		return "Critical feedback incoming! 🚨"
	case Important:
		return "Important feedback incoming! 🚨"
	case Warning:
		return "Warning! 🚨"
	default:
		return "Just a note! 📝"
	}
}
