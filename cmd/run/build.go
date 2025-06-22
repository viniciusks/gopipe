package run

import (
	"os"

	"github.com/spf13/cobra"
)

func checkoutProject(cmd *cobra.Command, cloneUrl string, branch string) {
	// This function can be used to checkout a project from a repository
	// For example, you can use git commands to clone or checkout a specific branch
	cmd.Println("Checking out the project...")
	cmd.Println("Clone URL:", cloneUrl)
	cmd.Println("Branch:", branch)
}

func identifyTechnology(cmd *cobra.Command, projectName string) string {
	// This function can be used to identify the technology used in the project
	// For example, you can check for specific files or configurations that indicate the technology
	// cmd.Println("Identifying technology...")
	// Here you can implement logic to identify the technology based on the project structure
	files, err := os.ReadDir("/tmp/gopipe/" + projectName)

	if err != nil {
		cmd.Println("Error reading directory:", err)
		return "unknown"
	}

	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == "go.mod" {
				cmd.Println("Technology identified: Go")
				return "go"
			} else if file.Name() == "requirements.txt" {
				cmd.Println("Technology identified: Python")
				return "python"
			} else if file.Name() == "package.json" {
				cmd.Println("Technology identified: Node.js")
				return "nodejs"
			} else if file.Name() == "pom.xml" {
				cmd.Println("Technology identified: Java (Maven)")
				return "javaMaven"
			}
		}
	}

	cmd.Println("Technology could not be identified, defaulting to 'unknown'")
	return "unknown"
}

var buildCmd = &cobra.Command {
	Use:  "build",
	Short: "Builds a project",
	Long:  "This command builds a project defined in the configuration file. It compiles the source code, runs tests, and prepares the project for deployment.",
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can implement the logic to build a project
		// For example, read a configuration file and execute the build steps defined in it
		// technology, _ := cmd.Flags().GetString("technology")
		checkout, _ := cmd.Flags().GetBool("checkout")
		cloneUrl, _ := cmd.Flags().GetString("cloneUrl")
		branch, _ := cmd.Flags().GetString("branch")
		projectName, _ := cmd.Flags().GetString("projectName")

		if checkout {
			// If checkout is true, we need to clone the project from the repository
			if cloneUrl == "" {
				cmd.Println("Error: --cloneUrl must be specified when --checkout is true")
				return
			}

			if branch == "" {
				branch = "main" // Default branch if not specified
				cmd.Println("No branch specified, defaulting to 'main'")
			}

			checkoutProject(cmd, cloneUrl, branch)
		}

		if projectName == "" && !checkout {
			cmd.Println("Error: --projectName must be specified")
			return
		}

		cmd.Printf("Project Name: %s\n", projectName)
		technology := identifyTechnology(cmd, projectName)

		cmd.Printf("Building project '%s' using technology '%s'\n", projectName, technology)

		cmd.Println("Build process completed successfully!")
		cmd.Println("You can now deploy the project or run it locally.")
		cmd.Println("For more information, refer to the documentation of the technology used.")
	},
}

func init() {
	// buildCmd.Flags().StringP("technology", "t", "", "Specify the technology to use for building the project (e.g., 'go', 'python', 'nodejs')")
	buildCmd.Flags().BoolP("checkout", "C", false, "Checkout the project before building")
	buildCmd.Flags().StringP("cloneUrl", "c", "", "Determine the repository URL to clone the project from")
	buildCmd.Flags().StringP("branch", "b", "", "Specify the branch to build from")
	buildCmd.Flags().StringP("projectName", "p", "", "Specify the name of the project to build")
	// Add the build command to the run command
	RunCmd.AddCommand(buildCmd)
}