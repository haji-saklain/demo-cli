package main

import (
	"errors"
	"log"
	"strings"

	"github.com/haji-saklain/commands-cli/commands"
	"github.com/spf13/cobra"
)

var helloFlag bool
var goodbyeFlag bool
var countFlag int
var timeFlag bool
var helloArgs string
var goodbyeArgs string

func main() {
	var rootCmd = &cobra.Command{
		Use:   "commands-cli",
		Short: "A simple CLI tool",
		Long:  `A simple CLI tool built with Cobra.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := validateFlags(); err != nil {
				log.Fatalf("Error: %v", err)
			}
			if helloFlag {
				if len(args) > 0 {
					log.Fatalf("Invalid command: %s", strings.Join(args, " "))
				}
				logIfError(commands.Hello(helloArgs))
			}
			if goodbyeFlag {
				if len(args) > 0 {
					log.Fatalf("Invalid command: %s", strings.Join(args, " "))
				}
				logIfError(commands.Goodbye(goodbyeArgs))
			}
			if countFlag > 0 {
				if countFlag > 15 {
					log.Fatalf("Error: value limit exceeded: maximum allowed is 15")
				}
				logIfError(commands.CountNumbers(countFlag))
			} else if countFlag < 0 {
				log.Fatalf("Error: count cannot be negative")
			}
			if timeFlag {
				logIfError(commands.PrintCurrentTime())
			}
		},
	}

	rootCmd.Flags().BoolVarP(&helloFlag, "hello", "H", false, "Print hello message")
	rootCmd.Flags().BoolVarP(&goodbyeFlag, "goodbye", "G", false, "Print goodbye message")
	rootCmd.Flags().IntVarP(&countFlag, "count", "C", 0, "Count numbers up to n")
	rootCmd.Flags().BoolVarP(&timeFlag, "time", "T", false, "Print current time")
	rootCmd.Flags().StringVar(&helloArgs, "hello-args", "", "Additional arguments for hello command")
	rootCmd.Flags().StringVar(&goodbyeArgs, "goodbye-args", "", "Additional arguments for goodbye command")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func logIfError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

func validateFlags() error {
	if helloFlag && goodbyeFlag {
		return errors.New("both hello and goodbye flags are not allowed together")
	}
	return nil
}
