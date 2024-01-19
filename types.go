package pplx

import (
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
)

// A type to encapsulate the names of the available models
type ModelType string

// The name of the model that will complete your prompt
const (
	Pplx7BCHat           ModelType = "pplx-7b-chat"
	Pplx70BChat          ModelType = "pplx-70b-chat"
	Pplx7BOnline         ModelType = "pplx-7b-online"
	Pplx70BOnline        ModelType = "pplx-70b-online"
	Llama270BChat        ModelType = "llama-2-70b-chat"
	CodeLlama34BInstruct ModelType = "codellama-34b-instruct"
	Mistral7BInstruct    ModelType = "mistral-7b-instruct"
	Mixtral8x7BInstruct  ModelType = "mixtral-8x7b-instruct"
)

var (
	ContextLength = map[ModelType]int{
		CodeLlama34BInstruct: 16384,
		Llama270BChat:        4096,
		Mistral7BInstruct:    4096,
		Mixtral8x7BInstruct:  4096,
		Pplx7BCHat:           8192,
		Pplx70BChat:          4096,
		Pplx7BOnline:         4096,
		Pplx70BOnline:        4096,
	}
)

const (
	API_URL             string = "api.perplexity.ai"
	API_CHAT_COMPLETION string = "/chat/completions"
)

// The role of the speaker in this turn of the conversation.
type SpeakerRole string

// The role of the speaker in this turn of the conversation.
const (
	RoleSystem    SpeakerRole = "system"
	RoleUser      SpeakerRole = "user"
	RoleAssistant SpeakerRole = "assistant"
)

var vld = validator.New()

// A message holds the information of a single chat input. Also holds context for who said it.
type ChatMessage struct {
	// The contents of the message in this turn of conversation.
	Content string `validate:"required" json:"content"`
	// The role of the speaker in this turn of conversation.
	// After the (optional) system message, user and assistant roles should alternate with user then assistant, ending in user.
	Role SpeakerRole `validate:"required" json:"role"`
}

type ChatCompletionRequest struct {
	// The name of the model that will complete your prompt
	Model ModelType `json:"model"`
	// A list of messages comprising the conversation so far.
	Messages []ChatMessage `validate:"required" json:"messages"`
	// The maximum number of completion tokens returned by the API.
	// The total number of tokens requested in max_tokens plus the number of prompt tokens sent in messages must not exceed the context window token limit of model requested.
	// If left unspecified, then the model will generate tokens until either it reaches its stop token or the end of its context window.
	MaxTokens *int `validate:"omitnil,gte=0" json:"max_tokens"`
	// The amount of randomness in the response, valued between 0 inclusive and 2 exclusive.
	// Higher values are more random, and lower values are more deterministic. You should either set temperature or TopP, but not both.
	Temperature *int `validate:"omitnil,gte=0,lte=1" json:"temperature"`
	// The nucleus sampling threshold, valued between 0 and 1 inclusive.
	// For each subsequent token, the model considers the results of the tokens with top_p probability mass. You should either alter temperature or top_p, but not both.
	TopP *int `validate:"omitnil,gte=0,lte=1" json:"top_p"`
	// The number of tokens to keep for highest top-k filtering, specified as an integer between 0 and 2048 inclusive.
	// If set to 0, top-k filtering is disabled.
	TopK int `validate:"gte=0,lte=2048" json:"top_k"`
	// Determines whether or not to incrementally stream the response with
	// [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#event_stream_format) with content-type: text/event-stream.
	Stream bool `json:"stream"`
	// A value between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics. Incompatible with frequency_penalty.
	PresencePenalty *int `validate:"omitnil,gte=-2,lte=2" json:"presence_penalty"`
	// A multiplicative penalty greater than 0. Values greater than 1.0 penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim. A value of 1.0 means no penalty. Incompatible with presence_penalty.
	FrequencyPenalty *int `validate:"omitnil,gt=0" json:"frequency_penalty"`
}

// Validates rules set by the API and returns a valid request struct if successful. Returns an error otherwise.
func newChatCompletionRequest(opts ChatCompletionOpts) (*ChatCompletionRequest, error) {
	if err := vld.Struct(opts); err != nil {
		return nil, err
	}

	// temperature and top_p are not compatible
	if opts.Temperature != nil && opts.TopP != nil {
		return nil, fmt.Errorf("cannot set both temperature and top_p")
	}

	// presence_penalty and frequency_penalty are not compatible
	if opts.FrequencyPenalty != nil && opts.PresencePenalty != nil {
		return nil, fmt.Errorf("cannot set both frequency_penalty and presence_penalty")
	}
	if opts.Stream {
		return nil, fmt.Errorf("stream not implemented")
	}

	if opts.Model == "" {
		// Set the default to mistral-7b-instruct
		opts.Model = string(Mistral7BInstruct)
	}

	return &ChatCompletionRequest{
		Model:            ModelType(opts.Model),
		Messages:         opts.Messages,
		MaxTokens:        opts.MaxTokens,
		Temperature:      opts.Temperature,
		TopP:             opts.TopP,
		TopK:             opts.TopK,
		Stream:           opts.Stream,
		PresencePenalty:  opts.PresencePenalty,
		FrequencyPenalty: opts.FrequencyPenalty,
	}, nil
}

// ChatCompletionResponse holds the returned response as a struct
type ChatCompletionResponse struct {
	// An ID generated uniquely for each response.
	Id string `json:"id"`
	// The model used to generate the response.
	Model string `json:"model"`
	// The Unix timestamp (in seconds) of when the completion was created.
	Created int `json:"created"`
	// The list of completion choices the model generated for the input prompt.
	Choices []struct {
		//
		Index int `json:"index"`
		// The reason the model stopped generating tokens.
		// Possible values include stop if the model hit a natural stopping point, or length if the maximum number of tokens specified in the request was reached.
		FinishReason string `json:"finish_reason"`
		// The message generated by the model.
		Message ChatMessage `json:"message"`
		// The incrementally streamed next tokens. Only meaningful when stream = true.
		Delta ChatMessage `json:"delta"`
	} `json:"choices"`
	// Usage statistics for the completion request.
	Usage struct {
		// The number of tokens provided in the request prompt.
		PromptTokens int `json:"prompt_tokens"`
		// The number of tokens generated in the response output.
		CompletionTokens int `json:"completion_tokens"`
		// The total number of tokens used in the chat completion (prompt + completion).
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}

// Default logger of the package
var Logger *slog.Logger = slog.Default()
