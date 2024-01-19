package pplx

import (
	"net/http"
)

var apikey string = ""

func WithAPIKey(s string) {
	apikey = s
}

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
