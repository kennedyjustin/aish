package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kennedyjustin/aish/openai"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	asciiArt = `
________  ___  ________  ___  ___
|\   __  \|\  \|\   ____\|\  \|\  \
\ \  \|\  \ \  \ \  \___|\ \  \\\  \
 \ \   __  \ \  \ \_____  \ \   __  \
  \ \  \ \  \ \  \|____|\  \ \  \ \  \
   \ \__\ \__\ \__\____\_\  \ \__\ \__\
    \|__|\|__|\|__|\_________\|__|\|__|
                  \|_________|

`
	commandName      = "aish"
	configFileSuffix = "json"
	defaultShell     = "/bin/bash"
)

var (
	configFile                       string
	requiredFieldsViperStringToLabel = map[string]string{
		"openai-secret-key": "OpenAI Secret Key",
	}
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", fmt.Sprintf("config file (default is $HOME/.%s.%s)", commandName, configFileSuffix))
}

var rootCmd = &cobra.Command{
	Use:   "aish",
	Short: "aish is an artificially intelligent shell.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(asciiArt)
		}

		if configFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(configFile)
		} else {
			// Find home directory.
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)

			// Search config in home directory with name ".aish.json".
			viper.AddConfigPath(home)
			viper.SetConfigType(configFileSuffix)
			configName := fmt.Sprintf(".%s.%s", commandName, configFileSuffix)
			viper.SetConfigName(configName)
			configFile = filepath.Join(home, configName)
		}

		viper.AutomaticEnv()

		// Ensure config is setup
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("Setting up aish...")
				fmt.Println("")

				// Get required fields for config file
				for requiredFieldViperString, requiredFieldLabel := range requiredFieldsViperStringToLabel {
					viper.SetDefault(requiredFieldViperString, promptUser(requiredFieldLabel))
				}
				fmt.Println("")

				// Write config file
				err := viper.WriteConfigAs(configFile)
				if err != nil {
					fmt.Printf("Error saving configuration to %s: %s\n", configFile, err.Error())
					os.Exit(1)
				} else {
					fmt.Printf("Successfully wrote config to %s\n", configFile)
				}

			} else {
				fmt.Println("Unknown error: ", err.Error())
				os.Exit(1)
			}

			teachUsage()
		} else if len(args) == 0 {
			teachUsage()
			return
		}

		shell := os.Getenv("SHELL")
		if len(shell) == 0 {
			shell = defaultShell
		}

		// Do Completion
		completedText, err := openai.CompleteText(strings.Join(args, " "), os.Getenv("SHELL"))
		if err != nil {
			fmt.Println("Error retrieving completion from OpenAI: ", err.Error())
			os.Exit(1)
		}

		// TODO: Implement better user experience than simply printing the output
		fmt.Println(completedText)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func promptUser(label string) string {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Failed to read %s: %v\n", label, err)
		os.Exit(1)
	}

	return result
}

func teachUsage() {
	// fmt.Println("TODO: TEACH HOW TO USE")
}
