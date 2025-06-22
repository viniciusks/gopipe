package run

import "github.com/spf13/cobra"

var buildCmd = &cobra.Command {
	Use:  "build",
	Short: "Builds a project",
	Long:  "This command builds a project defined in the configuration file. It compiles the source code, runs tests, and prepares the project for deployment.",
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can implement the logic to build a project
		// For example, read a configuration file and execute the build steps defined in it
		technology, _ := cmd.Flags().GetString("technology")
		cmd.Println("Building the project...")
		cmd.Printf("Using technology: %s\n", technology)
	},
}

func init() {
	buildCmd.Flags().StringP("technology", "t", "", "Specify the technology to use for building the project (e.g., 'go', 'python', 'nodejs')")
	// Add the build command to the run command
	RunCmd.AddCommand(buildCmd)
}