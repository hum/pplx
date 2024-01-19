# pplx

Wrapper for the [Perplexity API](https://blog.perplexity.ai/blog/introducing-pplx-api).

## Usage

```go
import "github.com/hum/pplx"

pplx.WithAPIKey("Bearer [YOUR_API_KEY]")
response, _ := pplx.ChatComplete(pplx.ChatCompletionOpts{
    Messages: []pplx.ChatMessage{{Content: "hello", Role: pplx.RoleUser}},
})
fmt.Println(r.Choices[0].Message.Content)
// Hello there! How can I assist you today? If you have any questions or need help with something, feel free to ask.
// ...
```

## Install

```bash
> go get "github.com/hum/pplx"
```
