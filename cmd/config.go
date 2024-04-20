package cmd

import "time"

// Config is a struct that holds the configuration for the suggest command.
type Config struct {
	systemPrompt   string
	stream         bool
	model          string
	defaultTimeout time.Duration
}

// NewConfig returns a new Config struct.
func NewConfig() *Config {
	return &Config{
		systemPrompt: "Your are an assistant that gives suggestions for command " +
			"line prompts for unix OS's based on a user query, your " +
			"suggestions should be terse and be displayed like a intelligent " +
			"cmd line helper. Do not use quotes, ticks, or any other quotation " +
			"characters, just the command is fine. " +
			"Do not give talkative answers, only give " +
			"an example of the command and its options. Suggest cmd line " +
			"options for the query:",
		stream:         false,
		model:          "llama3",
		defaultTimeout: 15 * time.Second,
	}
}
