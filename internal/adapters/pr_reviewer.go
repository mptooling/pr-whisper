package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mptooling/pr-whisper/internal/domain"
	"io"
	"net/http"
)

type PrClient interface {
	comment([]*domain.Comment) error
}

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

func (client PrReviewer) comment(comments []*domain.Comment) error {
	if len(comments) == 0 {
		fmt.Println("No comments to make")
		return nil
	}

	return client.commentWhispers(comments)
}

func (client PrReviewer) commentWhispers(comments []*domain.Comment) error {
	content := make(map[string][]domain.Comment)
	for _, comment := range comments {
		content[comment.WhisperName] = append(content[comment.WhisperName], *comment)
	}

	body := `<details>` + "\n\n" + `<summary>🤫 Psst... Here is a list of potential issues:</summary>` + "\n"
	for whisperName, commentList := range content {
		emoji := client.getEmojiForSection(commentList[0].Severity)
		body += fmt.Sprintf("\n%s %s\n", emoji, whisperName)
		for _, comment := range commentList {
			body += fmt.Sprintf("- [ ] Affected file `%s`. %s\n", comment.FilePath, comment.Content) // todo :: must be configurable
		}
	}

	body += "\n" + `</details>` + "\n"

	review := domain.PRReview{
		Body:  body,
		Event: "COMMENT",
	}

	jsonData, err := json.Marshal(review)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		fmt.Println("Review:", review)

		return err
	}

	return client.send(jsonData)
}

func (client PrReviewer) send(jsonData []byte) error {
	req, err := http.NewRequest("POST", client.url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		fmt.Println("Data:", jsonData)

		return err
	}

	for key, value := range client.headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending req:", err)
		fmt.Println("req:", req)
		fmt.Println("resp:", resp)

		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error sending req:", err)

		return err
	}

	fmt.Println(string(data))

	if resp.Status != "200 OK" {
		return fmt.Errorf("error commenting on PR: %s", resp.Status)
	}

	return nil
}

func (client PrReviewer) getEmojiForSection(severity int) string {
	switch severity {
	case domain.Important:
		return "🟣"
	case domain.Warning:
		return "🟠"
	case domain.Caution:
		return "🔴"
	default:
		return "🟢"
	}
}
