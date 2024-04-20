package cmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	ollama "github.com/cdpierse/commander/internal"
	"github.com/ollama/ollama/api"
)

var suggestFlag string

var rootCmd = &cobra.Command{
	Use:   "commander",
	Short: "commander is a CLI for getting command line copilot suggestions",
	Long:  "commander is a CLI for getting command line copilot suggestions.",
	Run: func(cmd *cobra.Command, args []string) {
		config := NewConfig()
		ctx, cancel := context.WithTimeout(context.Background(), config.defaultTimeout)
		defer cancel()
		request := &api.ChatRequest{
			Model: config.model,
			Messages: []api.Message{
				{
					Role:    "system",
					Content: config.systemPrompt,
				},
				{
					Role:    "user",
					Content: suggestFlag,
				},
			},
			Stream: &config.stream,
		}

		response, err := ollama.GetCommandSuggestion(ctx, request)
		if err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Request timed out")
				return
			}
			fmt.Println(err)
			return
		}
		green := color.New(color.FgGreen).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		fmt.Printf("%s: %s\n", blue("CMD"), green(response.Message.Content))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&suggestFlag, "query", "q", "", "The query to get cmd line suggestions for")
}

// Execute is the entry point for the CLI
func Execute() error {
	return rootCmd.Execute()
}
