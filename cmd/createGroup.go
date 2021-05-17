package cmd

import (
	"github.com/spf13/cobra"
)

var createGroupCmd = &cobra.Command{
	Use:   "create-group group",
	Short: "Create a group",
	Long: `Create a group.

Create the group resource and associated kustomization file`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createGroup(args[0], false)
	},
}

func init() {
	rootCmd.AddCommand(createGroupCmd)
}
