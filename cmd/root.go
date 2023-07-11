package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oj/parser"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oj",
	Short: "oj is a CLI utility that converts the output of linux cli to JSON.",
	Long: `OJ, which stands for Output JSON, is a robust Command Line Interface (CLI) utility that is designed to transform the output of various Linux commands into a structured JSON format. 
	The utility is capable of handling a wide array of Linux commands and can interpret different output formats by employing various parsers.
	Users can specify the appropriate parser for each command, ensuring that the conversion to JSON is both accurate and reliable.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get parser from flag, or use default if not specified
		parserName, _ := cmd.Flags().GetString("parser")
		p, ok := parser.Parsers[parserName]
		if !ok {
			p = parser.Parsers["default"]
		}

		// Read from stdin
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("error reading input: %w", err)
		}

		// Parse input
		parsedOutput, err := p.Parse(string(input))
		if err != nil {
			return fmt.Errorf("error parsing input: %w", err)
		}

		// Convert to JSON and print
		jsonOutput, err := json.Marshal(parsedOutput)
		if err != nil {
			return fmt.Errorf("error converting to JSON: %w", err)
		}
		fmt.Println(string(jsonOutput))
		return nil
	},
}

func init() {
	parserNames := make([]string, 0, len(parser.Parsers))
	for name := range parser.Parsers {
		parserNames = append(parserNames, name)
	}
	rootCmd.Flags().StringP("parser", "p", "default", "specify the parser to use ("+strings.Join(parserNames, ", ")+")")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
