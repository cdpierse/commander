package ollama

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

// GetCommandSuggestion sends a chat request to the Ollama API and returns the response.
func GetCommandSuggestion(ctx context.Context, request *api.ChatRequest) (*api.ChatResponse, error) {
	// Create a client from environment settings
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("Failed to create client: %w", err)
	}

	// Define a function to handle the response from the chat API
	var chatResponse api.ChatResponse
	handleResponse := func(response api.ChatResponse) error {
		chatResponse = response
		return nil
	}
	// Send the chat request and handle the response
	err = client.Chat(ctx, request, handleResponse)
	if err != nil {
		return nil, fmt.Errorf("Chat request failed, are you sure Ollama is running? Err: %w", err)
	}

	return &chatResponse, nil
}
