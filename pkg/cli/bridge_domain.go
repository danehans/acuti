package cli

import (
	"github.com/spf13/cobra"
)

// brCmd represents the bridge-domain command
var brCmd = &cobra.Command{
	Use:   "bridge-domain",
	Short: "Manage bridge domains",
	Long:  `List, create & delete bridge domains.`,
}

func init() {
	RootCmd.AddCommand(brCmd)
}
