package cli

import (
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage ACI nodes",
	Long:  `List, create & delete ACI nodes.`,
}

func init() {
	RootCmd.AddCommand(nodeCmd)
}
