# pplx

Wrapper for the [Perplexity API](https://blog.perplexity.ai/blog/introducing-pplx-api).

## Usage

### Basic usage

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

### Tweak the parameters

```go
response, _ := pplx.ChatComplete(pplx.ChatCompletionOpts{
    Messages:    []pplx.ChatMessage{{Content: "hello", Role: pplx.RoleUser}},
    Model:       pplx.Mixtral8x7BInstruct,
    MaxTokens:   pplx.IntVar(50),
    Temperature: pplx.IntVar(1),
    
})
fmt.Println(r.Choices[0].Message.Content)
// Hello! It's nice to meet you.
// Is there something you'd like to talk about or ask me?
// I'm here to help with information and answer any questions you might have
// to the best of my ability. I can assist with
fmt.Println(r.Usage.CompletionTokens)
// 50
```


## Install

```bash
> go get "github.com/hum/pplx"
```
