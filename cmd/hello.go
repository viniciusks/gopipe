package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Exibe uma mensagem de saudação",
	Long:  `Este comando exibe uma mensagem de saudação personalizada.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "Mundo"
		}
		fmt.Printf("Olá, %s!\n", name)
	},
}

func init() {
	// Adiciona a flag "name" ao comando hello
	helloCmd.Flags().StringP("name", "n", "", "Nome da pessoa a ser saudada")
	// Registra o comando hello no rootCmd
	rootCmd.AddCommand(helloCmd)
}