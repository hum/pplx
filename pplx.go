package pplx

import (
	"net/http"
)

var apikey string = ""

// Sets the API bearer token for the whole package to use.
func WithAPIKey(s string) {
	apikey = s
}

// The main entrypoint to the Perplexity API which accepts a set of parameters as opts
// and returns the successful response in a struct. Otherwise returns an error.
func ChatComplete(opts ChatCompletionOpts) (*ChatCompletionResponse, error) {
	r, err := newChatCompletionRequest(opts)
	if err != nil {
		return nil, err
	}
	req, err := buildRequest(http.MethodPost, API_CHAT_COMPLETION, apikey, r)
	if err != nil {
		return nil, err
	}
	return performRequest[ChatCompletionResponse](req)
}

// Helper function to transform an integer value into a pointer.
func IntVar(i int) *int {
	return &i
}
