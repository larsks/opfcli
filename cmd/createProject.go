package cmd

import (
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "create-project projectName projectOwner",
	Short: "Onboard a new project into Operate First",
	Long: `Onboard a new project into Operate First.

- Register a new group
- Register a new namespace with appropriate role bindings for your group
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectOwner := args[1]
		projectDescription := cmd.Flag("description").Value.String()

		createNamespace(projectName, projectOwner, projectDescription)
		createAdminRoleBinding(projectName, projectOwner)
		createGroup(projectOwner, true)
	},
}

func init() {
	rootCmd.AddCommand(createProjectCmd)

	createProjectCmd.PersistentFlags().StringP(
		"description", "d", "", "Team description")
}
