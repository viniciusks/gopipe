package run

import (
	"github.com/spf13/cobra"
)

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Runs a pipeline",
	Long:  "This command runs a pipeline defined in the configuration file. It executes the steps in the pipeline sequentially or in parallel as defined.",
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can implement the logic to run a pipeline
		// For example, read a configuration file and execute the defined steps
		cmd.Println("Running the pipeline...")
	},
}

func init() {
	// Add the pipeline command to the run command
	RunCmd.AddCommand(pipelineCmd)
}