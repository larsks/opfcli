package cmd

import (
	"github.com/spf13/cobra"
)

var grantAccessCommand = &cobra.Command{
	Use:   "grant-access namespace group role",
	Short: "Grant a group access to a namespace",
	Long: `Grant a group access to a namespace.

Grant a group access to a namespace with the specifed role
(admin, edit, or view).`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		addGroupRBAC(args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(grantAccessCommand)
}
