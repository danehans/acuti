package cli

import (
	"github.com/spf13/cobra"
)

// vpoolCmd represents the vlan-pool command
var vpoolCmd = &cobra.Command{
	Use:   "vlan-pool",
	Short: "Manage vlan pools",
	Long:  `List, create & delete vlan pools.`,
}

func init() {
	RootCmd.AddCommand(vpoolCmd)
}
