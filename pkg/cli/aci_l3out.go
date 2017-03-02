package cli

import (
	"github.com/spf13/cobra"
)

// l3outCmd represents the l3-out command
var l3outCmd = &cobra.Command{
	Use:   "l3-out",
	Short: "Manage ACI Tenant Layer 3 Outside Network Connections",
	Long:  `List, create & delete ACI Tenant Layer 3 Outside Network Connections.`,
}

func init() {
	aciCmd.AddCommand(l3outCmd)
}
