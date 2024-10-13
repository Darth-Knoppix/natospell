/*
Copyright © 2024 Seth Corker hello@sethcorker.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "natospell [string to transform]",
	Short: "Transform a string into clear-code words using Radiotelephony Spelling Alphabet",
	Long:  `Transform a string using Radiotelephony Spelling Alphabet (NATO phonetic alphabet) so it can be read easily over telephone systems.`,

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var transformedBuilder strings.Builder
		var inputBuilder strings.Builder

		natoMap := map[string]string{
			"a": "Alpha",
			"b": "Bravo",
			"c": "Charlie",
			"d": "Delta",
			"e": "Echo",
			"f": "Foxtrot",
			"g": "Golf",
			"h": "Hotel",
			"i": "India",
			"j": "Juliett",
			"k": "Kilo",
			"l": "Lima",
			"m": "Mike",
			"n": "November",
			"o": "Oscar",
			"p": "Papa",
			"q": "Quebec",
			"r": "Romeo",
			"s": "Sierra",
			"t": "Tango",
			"u": "Uniform",
			"v": "Victor",
			"w": "Whiskey",
			"x": "Xray",
			"y": "Yankee",
			"z": "Zulu",
		}

		input := strings.ToLower(strings.TrimSpace(strings.Join(args, "")))

		for index, letter := range input {
			strLetter := string(letter)
			codeWord := natoMap[strLetter]

			if len(codeWord) > 0 {
				transformedBuilder.WriteString(codeWord)
				inputBuilder.WriteString(strLetter + strings.Repeat(" ", len(codeWord)-1))
			} else {
				transformedBuilder.WriteString(strLetter)
				inputBuilder.WriteString(strLetter)
			}

			if index+1 < len(input) {
				transformedBuilder.WriteString(" ")
				inputBuilder.WriteString(" ")
			}
		}

		fmt.Println("   Original: " + inputBuilder.String())
		fmt.Println("Transformed: " + transformedBuilder.String())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.natospell.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
