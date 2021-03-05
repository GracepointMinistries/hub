package admin

import (
	"github.com/spf13/cobra"
)

// Register adds all of the subcommands to the parent
func Register(parent *cobra.Command) {
	parent.AddCommand(impersonateCmd)
	parent.AddCommand(groupCmd)
	parent.AddCommand(userCmd)
	parent.AddCommand(syncCmd)
	parent.AddCommand(exportCmd)
}
