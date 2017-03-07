package cli

import (
	"github.com/spf13/cobra"
)

// l3outCmd represents the l3-out command
var l3outCmd = &cobra.Command{
	Use:   "l3-out",
	Short: "Manage Tenant Layer 3 Outside Network Connections",
	Long:  `List, create & delete Tenant Layer 3 Outside Network Connections.`,
}

func init() {
	RootCmd.AddCommand(l3outCmd)
}
