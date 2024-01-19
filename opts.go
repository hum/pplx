package pplx

// ChatCompletionOpts encapsulates the options provided by API.
type ChatCompletionOpts struct {
	// The name of the model that will complete your prompt
	Model string
	// A list of messages comprising the conversation so far.
	Messages []ChatMessage `validate:"required"`
	// The maximum number of completion tokens returned by the API.
	// The total number of tokens requested in max_tokens plus the number of prompt tokens sent in messages must not exceed the context window token limit of model requested.
	// If left unspecified, then the model will generate tokens until either it reaches its stop token or the end of its context window.
	MaxTokens *int `validate:"omitnil,gte=0"`
	// The amount of randomness in the response, valued between 0 inclusive and 2 exclusive.
	// Higher values are more random, and lower values are more deterministic. You should either set temperature or TopP, but not both.
	Temperature *int `validate:"omitnil,gte=0,lte=1"`
	// The nucleus sampling threshold, valued between 0 and 1 inclusive.
	// For each subsequent token, the model considers the results of the tokens with top_p probability mass. You should either alter temperature or top_p, but not both.
	TopP *int `validate:"omitnil,gte=0,lte=1"`
	// The number of tokens to keep for highest top-k filtering, specified as an integer between 0 and 2048 inclusive.
	// If set to 0, top-k filtering is disabled.
	TopK int `validate:"gte=0,lte=2048"`
	// Determines whether or not to incrementally stream the response with
	// [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#event_stream_format) with content-type: text/event-stream.
	Stream bool
	// A value between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics. Incompatible with frequency_penalty.
	PresencePenalty *int `validate:"omitnil,gte=-2,lte=2"`
	// A multiplicative penalty greater than 0. Values greater than 1.0 penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim. A value of 1.0 means no penalty. Incompatible with presence_penalty.
	FrequencyPenalty *int `validate:"omitnil,gt=0"`
}
