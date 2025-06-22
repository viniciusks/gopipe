package run

import (
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a pipeline or a task",
	Long:  "This command runs a pipeline or a specific task defined in the configuration file.",
	Run: func(cmd *cobra.Command, args []string) {
		// Aqui você pode implementar a lógica para executar um pipeline ou tarefa
		// Por exemplo, ler um arquivo de configuração e executar os passos definidos nele
		cmd.Println("Running the pipeline or task...")
	},
}
