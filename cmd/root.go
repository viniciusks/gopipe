package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
    tag    = "dev"
    commit = "none"
    date   = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "gopipe",
	Short: "gopipe is a CLI tool for something",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gopipe! Use 'gopipe --help' to see available commands.")
	},
}

func SetVersionVars(t, c, d string) {
    tag = t
    commit = c
    date = d
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}