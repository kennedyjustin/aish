package openai

import (
	"context"
	"fmt"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/viper"
)

const promptTemplate = `
#!%s
# %s
`

func CompleteText(input string, shell string) (string, error) {
	c := gogpt.NewClient(viper.GetString("openai-secret-key"))
	resp, err := c.CreateCompletion(context.Background(), gogpt.CompletionRequest{
		Model:       gogpt.CodexCodeDavinci002,
		MaxTokens:   30,
		Temperature: 0.1,
		Prompt:      fmt.Sprintf(promptTemplate, shell, input),
	})
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.Split(resp.Choices[0].Text, "</code>")[0]), nil
}
