package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		config := NewConfig()
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
		// Replace backticks with single quotes
		output := strings.ReplaceAll(response.Message.Content, "`", "")
		green := color.New(color.FgGreen).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		fmt.Printf("%s: %s\n", blue("CMD"), green(output))
		// fmt.Printf("CMD: %s\n", green(output))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&suggestFlag, "query", "q", "", "The query to get cmd line suggestions for")
}

// Execute is the entry point for the CLI
func Execute() error {
	return rootCmd.Execute()
}
