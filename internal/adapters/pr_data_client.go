package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/mptooling/pr-whisper/internal/domain"
	"io"
	"net/http"
)

type prFilesClient struct {
	request *http.Request
}

func NewPrDataClient(apiUrl string, token string, repo string, pullRequestNumber string) PrFilesClient {
	url := fmt.Sprintf("%s/repos/%s/pulls/%s/files", apiUrl, repo, pullRequestNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	return &prFilesClient{
		request: req,
	}
}

func (client prFilesClient) GetPrFiles() (domain.DiffEntries, error) {
	resp, err := http.DefaultClient.Do(client.request)
	fmt.Println("Request:", client.request)
	fmt.Println("Response:", resp)

	if err != nil {
		fmt.Println("Error getting PR files:", err)

		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var files domain.DiffEntries
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, err
	}

	return files, nil
}
