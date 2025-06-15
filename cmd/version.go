package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the application version",
	Long:  `This command displays the current version of the gopipe application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("GoPipe version: %s\nCommit: %s\nDate: %s\n", tag, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}