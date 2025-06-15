package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Displays a greeting message",
	Long:  `This command displays a personalized greeting message.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	// Adiciona a flag "name" ao comando hello
	helloCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	// Registra o comando hello no rootCmd
	rootCmd.AddCommand(helloCmd)
}