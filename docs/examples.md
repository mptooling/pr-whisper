## PR FILES response example
URL: https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#create-a-review-for-a-pull-request
<details>
<summary>Content</summary>

```json
[
  {
    "sha": "cbcdc0c8d0fd864c1d87e941d1f0dba629f98c9a",
    "filename": ".github/workflows/ci.yml",
    "status": "modified",
    "additions": 2,
    "deletions": 1,
    "changes": 3,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/.github%2Fworkflows%2Fci.yml",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/.github%2Fworkflows%2Fci.yml",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/.github%2Fworkflows%2Fci.yml?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d",
    "patch": "@@ -15,7 +15,8 @@ jobs:\n       - name: PR Whisper\n         uses: ./\n         env:\n-          GITHUB_TOKEN: ${{ secrets.PAT }}\n+          GITHUB_TOKEN: ${{ github.token }}\n+          GH_AUTH_TOKEN: ${{ secrets.PAT }}\n           GITHUB_PULL_REQUEST_NUMBER: ${{ github.event.number }}\n           GITHUB_SHA: ${{ github.sha }}\n           GITHUB_REPOSITORY: ${{ github.repository }}"
  },
  {
    "sha": "7292881bf8dc7d134ffa8a2bdd76ca53fbb422bc",
    "filename": "Dockerfile",
    "status": "modified",
    "additions": 16,
    "deletions": 1,
    "changes": 17,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/Dockerfile",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/Dockerfile",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/Dockerfile?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d",
    "patch": "@@ -1,5 +1,20 @@\n FROM golang:1.23\n \n+ARG GITHUB_TOKEN=${GITHUB_TOKEN}\n+ENV GITHUB_TOKEN=${GITHUB_TOKEN}\n+\n+ARG GH_AUTH_TOKEN=${GH_AUTH_TOKEN}\n+ENV GH_AUTH_TOKEN=${GH_AUTH_TOKEN}\n+\n+ARG GITHUB_OWNER=${GITHUB_OWNER}\n+ENV GITHUB_OWNER=${GITHUB_OWNER}\n+\n+ARG GITHUB_REPOSITORY=${GITHUB_REPOSITORY}\n+ENV GITHUB_REPOSITORY=${GITHUB_REPOSITORY}\n+\n+ARG GITHUB_PULL_REQUEST_NUMBER=${GITHUB_PULL_REQUEST_NUMBER}\n+ENV GITHUB_PULL_REQUEST_NUMBER=${GITHUB_PULL_REQUEST_NUMBER}\n+\n WORKDIR /app\n \n COPY go.mod go.sum ./\n@@ -9,7 +24,7 @@ COPY *.go ./\n COPY entrypoint.sh /entrypoint.sh\n \n RUN #CGO_ENABLED=0 GOOS=linux go build -o /prwhisper\n-RUN go build -o /prwhisper\n+RUN go build -o /app/prwhisper\n \n RUN chmod +x /entrypoint.sh\n "
  },
  {
    "sha": "ee7dc752b5fdf01c11849a15c723d7adbc4a23be",
    "filename": "README.md",
    "status": "added",
    "additions": 93,
    "deletions": 0,
    "changes": 93,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/README.md",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/README.md",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/README.md?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d",
    "patch": "@@ -0,0 +1,93 @@\n+# GitHub PR Whisperer\n+\n+This project is a Go application designed to interact with the GitHub API to whisper to the PR if it contains rules violations.\n+\n+Moto: \"Whisper to the PR, don't shout!\"\n+\n+## Environment Variables\n+\n+The following environment variables need to be set for the application to function correctly:\n+\n+- `GITHUB_TOKEN`: Autogenerated GitHub token that can be taken from the action.\n+- `GH_AUTH_TOKEN`: Your GitHub authentication token. Should be a personal access token with the `repo` scope.\n+- `GITHUB_REPOSITORY`: The name of the repository.\n+- `GITHUB_PULL_REQUEST_NUMBER`: The pull request number.\n+\n+## Usage\n+\n+### Running Locally\n+\n+1. Clone the repository:\n+    ```sh\n+    git clone httpshttps://api.github.com/mptooling/pr-whisper.git\n+    cd pr-whisper\n+    ```\n+\n+2. Set the required environment variables:\n+    ```sh\n+    export GITHUB_TOKEN=your_github_token\n+    export GH_AUTH_TOKEN=github_auth_token\n+    export GITHUB_OWNER=your_github_owner\n+    export GITHUB_REPOSITORY=your_github_repository\n+    export GITHUB_PULL_REQUEST_NUMBER=your_pull_request_number\n+    ```\n+\n+3. Build and run the application:\n+    ```sh\n+    go build -o prwhisperer\n+    ./prwhisperer\n+    ```\n+\n+### Running with Docker\n+\n+1. Build the Docker image:\n+    ```sh\n+    docker build -t prwhisperer .\n+    ```\n+\n+2. Run the Docker container:\n+    ```sh\n+    docker run -e GITHUB_TOKEN=your_github_token \\\n+               -e GH_AUTH_TOKEN=your_github_auth_token \\\n+               -e GITHUB_OWNER=your_github_owner \\\n+               -e GITHUB_REPOSITORY=your_github_repository \\\n+               -e GITHUB_PULL_REQUEST_NUMBER=your_pull_request_number \\\n+               prwhisperer\n+    ```\n+\n+### Using as a GitHub Action\n+\n+1. Create a `.github/workflows/pr-whisperer.yml` file in your repository with the following content:\n+    ```yaml\n+    name: PR Whisperer\n+\n+    on:\n+      pull_request:\n+        types: [opened, synchronize, reopened]\n+\n+    jobs:\n+      pr-whisperer:\n+        runs-on: ubuntu-latest\n+\n+        steps:\n+        - name: Checkout repository\n+          uses: actions/checkout@v2\n+\n+        - name: Set up Go\n+          uses: actions/setup-go@v2\n+          with:\n+            go-version: 1.23\n+\n+        - name: Build and run PR Whisperer\n+          uses: mptooling/pr-whisper@main # Not released yet\n+          env:\n+            GITHUB_TOKEN: ${{ github.token }}\n+            GH_AUTH_TOKEN: ${{ secrets.PAT }} # Requires secret named PAT\n+            GITHUB_SHA: ${{ github.sha }}\n+            GITHUB_REPOSITORY: ${{ github.repository }}\n+            GITHUB_PULL_REQUEST_NUMBER: ${{ github.event.number }}\n+    ```\n+\n+## License\n+\n+This project is licensed under the MIT License.\n\\ No newline at end of file"
  },
  {
    "sha": "5e579a893cc806e3b69cb22ab81ea79a2f170725",
    "filename": "adapters/interfaces.go",
    "status": "removed",
    "additions": 0,
    "deletions": 7,
    "changes": 7,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/456ad4995284c44a47ec13bfac616b4cdcaa658a/adapters%2Finterfaces.go",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/456ad4995284c44a47ec13bfac616b4cdcaa658a/adapters%2Finterfaces.go",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/adapters%2Finterfaces.go?ref=456ad4995284c44a47ec13bfac616b4cdcaa658a",
    "patch": "@@ -1,7 +0,0 @@\n-package adapters\n-\n-import \"net/http\"\n-\n-type GithubClient interface {\n-\tHttpCall()(*http.Response, error)\n-}"
  },
  {
    "sha": "8dbc4e5e70544544cea8e7be15bb4fdca61df95c",
    "filename": "cmd/app/main.go",
    "status": "removed",
    "additions": 0,
    "deletions": 81,
    "changes": 81,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/456ad4995284c44a47ec13bfac616b4cdcaa658a/cmd%2Fapp%2Fmain.go",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/456ad4995284c44a47ec13bfac616b4cdcaa658a/cmd%2Fapp%2Fmain.go",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/cmd%2Fapp%2Fmain.go?ref=456ad4995284c44a47ec13bfac616b4cdcaa658a",
    "patch": "@@ -1,81 +0,0 @@\n-package main\n-\n-import (\n-\t\"bytes\"\n-\t\"encoding/json\"\n-\t\"fmt\"\n-\t\"io\"\n-\t\"net/http\"\n-\t\"os\"\n-)\n-\n-func getPRFiles(owner, repo, pullNumber, authHeader string) ([]map[string]interface{}, error) {\n-\tf := NewPRFilesClient(\"https://api.github.com\", authHeader, owner, repo, pullNumber)\n-\n-\tresp, err := prFilesClient.HttpCall()\n-\tif err != nil {\n-\t\treturn nil, err\n-\t}\n-\tdefer resp.Body.Close()\n-\n-\tbody, err := io.ReadAll(resp.Body)\n-\tif err != nil {\n-\t\treturn nil, err\n-\t}\n-\n-\tvar files []map[string]interface{}\n-\tif err := json.Unmarshal(body, &files); err != nil {\n-\t\treturn nil, err\n-\t}\n-\n-\treturn files, nil\n-}\n-\n-func main() {\n-\ttoken := os.Getenv(\"GITHUB_TOKEN\")\n-\tif token == \"\" {\n-\t\tfmt.Println(\"Error: GITHUB_TOKEN environment variable is not set.\")\n-\t\treturn\n-\t}\n-\n-\tauthHeader := \"Bearer \" + token\n-\n-\tfiles, err := getPRFiles(\"mptooling\", \"pr-whisper\", \"1\", authHeader)\n-\tif err != nil {\n-\t\tfmt.Println(\"Error fetching PR files:\", err)\n-\t\treturn\n-\t}\n-\tfor _, file := range files {\n-\t\tfmt.Println(file[\"filename\"])\n-\t}\n-\n-\turl := \"https://api.github.com/repos/mptooling/pr-whisper/pulls/1/reviews\"\n-\tjsonData := `{\"body\": \"This is close to perfect! Please address the suggested inline change.\",\"event\": \"COMMENT\"}`\n-\n-\treq, err := http.NewRequest(\"POST\", url, bytes.NewBuffer([]byte(jsonData)))\n-\tif err != nil {\n-\t\tfmt.Println(\"Error creating request:\", err)\n-\t\treturn\n-\t}\n-\n-\treq.Header.Set(\"Accept\", \"application/vnd.github+json\")\n-\treq.Header.Set(\"Authorization\", authHeader)\n-\treq.Header.Set(\"X-GitHub-Api-Version\", \"2022-11-28\")\n-\treq.Header.Set(\"Content-Type\", \"application/json\")\n-\n-\tresp, err := http.DefaultClient.Do(req)\n-\tif err != nil {\n-\t\tfmt.Println(\"Error making request:\", err)\n-\t\treturn\n-\t}\n-\tdefer resp.Body.Close()\n-\n-\tbody, err := io.ReadAll(resp.Body)\n-\tif err != nil {\n-\t\tfmt.Println(\"Error reading response:\", err)\n-\t\treturn\n-\t}\n-\n-\tfmt.Println(\"Response Status:\", resp.Status)\n-\tfmt.Println(\"Response Body:\", string(body))\n-}"
  },
  {
    "sha": "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391",
    "filename": "go.sum",
    "status": "added",
    "additions": 0,
    "deletions": 0,
    "changes": 0,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/go.sum",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/go.sum",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/go.sum?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d"
  },
  {
    "sha": "5b8eab471a0449ea290d52ee12cbca46ea2333e8",
    "filename": "main.go",
    "status": "added",
    "additions": 62,
    "deletions": 0,
    "changes": 62,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/main.go",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/main.go",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/main.go?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d",
    "patch": "@@ -0,0 +1,62 @@\n+package main\n+\n+import (\n+\t\"encoding/json\"\n+\t\"fmt\"\n+\t\"io\"\n+\t\"os\"\n+)\n+\n+func getPRFiles() (DiffEntries, error) {\n+\ttoken := os.Getenv(\"GITHUB_TOKEN\")\n+\trepo := os.Getenv(\"GITHUB_REPOSITORY\")\n+\tpullNumber := os.Getenv(\"GITHUB_PULL_REQUEST_NUMBER\")\n+\tclient := NewPrFilesClient(\"https://api.github.com\", token, repo, pullNumber)\n+\n+\tresp, err := client.getPrFiles()\n+\tif err != nil {\n+\t\treturn nil, err\n+\t}\n+\tdefer resp.Body.Close()\n+\n+\tbody, err := io.ReadAll(resp.Body)\n+\tif err != nil {\n+\t\treturn nil, err\n+\t}\n+\n+\tvar files DiffEntries\n+\tif err := json.Unmarshal(body, &files); err != nil {\n+\t\treturn nil, err\n+\t}\n+\n+\treturn files, nil\n+}\n+\n+func comment(message string) error {\n+\ttoken := os.Getenv(\"GH_AUTH_TOKEN\")\n+\trepo := os.Getenv(\"GITHUB_REPOSITORY\")\n+\tpullNumber := os.Getenv(\"GITHUB_PULL_REQUEST_NUMBER\")\n+\treviewer := NewPrReviewer(\"https://api.github.com\", token, repo, pullNumber)\n+\terr := reviewer.comment(message)\n+\tif err != nil {\n+\t\treturn err\n+\t}\n+\n+\treturn nil\n+}\n+\n+func main() {\n+\tfiles, err := getPRFiles()\n+\tif err != nil {\n+\t\tfmt.Println(\"Error getting PR files:\", err)\n+\t}\n+\n+\tfor _, file := range files {\n+\t\tfmt.Printf(\"File: %s. Status: %s \\n\", file.Filename, file.Status)\n+\t}\n+\n+\terr = comment(\"Hello from Go!\")\n+\tif err != nil {\n+\t\tfmt.Println(\"Error commenting on PR:\", err)\n+\t}\n+}"
  },
  {
    "sha": "b5883165f659d188da9e8cfb9503bb47ae9a5117",
    "filename": "models.go",
    "status": "added",
    "additions": 17,
    "deletions": 0,
    "changes": 17,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/models.go",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/models.go",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/models.go?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d",
    "patch": "@@ -0,0 +1,17 @@\n+package main\n+\n+type DiffEntry struct {\n+\tSha              string `json:\"sha\"`\n+\tFilename         string `json:\"filename\"`\n+\tStatus           string `json:\"status\"`\n+\tAdditions        int    `json:\"additions\"`\n+\tDeletions        int    `json:\"deletions\"`\n+\tChanges          int    `json:\"changes\"`\n+\tBlobURL          string `json:\"blob_url\"`\n+\tRawURL           string `json:\"raw_url\"`\n+\tContentsURL      string `json:\"contents_url\"`\n+\tPatch            string `json:\"patch,omitempty\"`\n+\tPreviousFilename string `json:\"previous_filename,omitempty\"`\n+}\n+\n+type DiffEntries []DiffEntry"
  },
  {
    "sha": "2174e3afb3c6629b05a3f289a3483f23931bdf5d",
    "filename": "pr_files_client.go",
    "status": "renamed",
    "additions": 4,
    "deletions": 4,
    "changes": 8,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/pr_files_client.go",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/pr_files_client.go",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/pr_files_client.go?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d",
    "patch": "@@ -1,4 +1,4 @@\n-package adapters\n+package main\n \n import (\n \t\"fmt\"\n@@ -9,8 +9,8 @@ type PrFilesClient struct {\n \trequest *http.Request\n }\n \n-func NewPrFilesClient(apiUrl string, token string, owner string, repo string, pullRequestNumber string) GithubClient {\n-\turl := fmt.Sprintf(\"%s/repos/%s/%s/pulls/%s/files\", apiUrl, owner, repo, pullRequestNumber)\n+func NewPrFilesClient(apiUrl string, token string, repo string, pullRequestNumber string) *PrFilesClient {\n+\turl := fmt.Sprintf(\"%s/repos/%s/pulls/%s/files\", apiUrl, repo, pullRequestNumber)\n \treq, err := http.NewRequest(\"GET\", url, nil)\n \tif err != nil {\n \t\tpanic(err)\n@@ -25,7 +25,7 @@ func NewPrFilesClient(apiUrl string, token string, owner string, repo string, pu\n \t}\n }\n \n-func (client PrFilesClient) HttpCall() (*http.Response, error) {\n+func (client PrFilesClient) getPrFiles() (*http.Response, error) {\n \tresp, err := http.DefaultClient.Do(client.request)\n \tif err != nil {\n \t\treturn nil, err",
    "previous_filename": "adapters/pr_files_client.go"
  },
  {
    "sha": "1d33990b9d12424d36a89f6dce2021513e1c562e",
    "filename": "pr_reviewer.go",
    "status": "added",
    "additions": 52,
    "deletions": 0,
    "changes": 52,
    "blob_url": "https://github.com/mptooling/pr-whisper/blob/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/pr_reviewer.go",
    "raw_url": "https://github.com/mptooling/pr-whisper/raw/244bc34f09b7ad93ab510e0ca4675aa20eb2750d/pr_reviewer.go",
    "contents_url": "https://api.github.com/repos/mptooling/pr-whisper/contents/pr_reviewer.go?ref=244bc34f09b7ad93ab510e0ca4675aa20eb2750d",
    "patch": "@@ -0,0 +1,52 @@\n+package main\n+\n+import (\n+\t\"bytes\"\n+\t\"fmt\"\n+\t\"io\"\n+\t\"net/http\"\n+)\n+\n+type PrReviewer struct {\n+\turl     string\n+\theaders map[string]string\n+}\n+\n+func NewPrReviewer(apiUrl string, token string, repo string, pullRequestNumber string) *PrReviewer {\n+\turl := fmt.Sprintf(\"%s/repos/%s/pulls/%s/reviews\", apiUrl, repo, pullRequestNumber)\n+\theaders := map[string]string{\n+\t\t\"Accept\":               \"application/vnd.github+json\",\n+\t\t\"Authorization\":        \"Bearer \" + token,\n+\t\t\"X-GitHub-Api-Version\": \"2022-11-28\",\n+\t}\n+\n+\treturn &PrReviewer{\n+\t\turl:     url,\n+\t\theaders: headers,\n+\t}\n+}\n+\n+func (client PrReviewer) comment(message string) error {\n+\tjsonData := `{\"body\":\"` + message + `\",\"event\": \"COMMENT\"}`\n+\n+\treq, err := http.NewRequest(\"POST\", client.url, bytes.NewBuffer([]byte(jsonData)))\n+\tif err != nil {\n+\t\tpanic(err)\n+\t}\n+\n+\tfor key, value := range client.headers {\n+\t\treq.Header.Set(key, value)\n+\t}\n+\n+\tresp, err := http.DefaultClient.Do(req)\n+\tif err != nil {\n+\t\treturn err\n+\t}\n+\tdefer resp.Body.Close()\n+\n+\t_, err = io.ReadAll(resp.Body)\n+\tif err != nil {\n+\t\treturn err\n+\t}\n+\treturn nil\n+}"
  }
]
```
</details>