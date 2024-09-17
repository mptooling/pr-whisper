package adapters

import "net/http"

type GithubClient interface {
	HttpCall()(*http.Response, error)
}
