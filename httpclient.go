package pplx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var httpclient *http.Client = &http.Client{Timeout: 60 * time.Second}

func buildRequest(httpMethod, endpoint, bearer string, body any) (*http.Request, error) {
	u := url.URL{
		Scheme: "https",
		Host:   API_URL,
		Path:   endpoint,
	}
	bbody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest(httpMethod, u.String(), bytes.NewReader(bbody))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", bearer)
	return r, nil
}

func performRequest[T any](r *http.Request) (*T, error) {
	resp, err := httpclient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non-successful request, status=%s body=%s", resp.Status, string(body))
	}

	result := new(T)
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if err = vld.Struct(result); err != nil {
		return nil, err
	}
	return result, nil
}
