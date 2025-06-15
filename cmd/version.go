package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0.0" // Version of the application

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Exibe a versão do aplicativo",
	Long:  `Este comando exibe a versão atual do aplicativo gopipe.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gopipe versão %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}