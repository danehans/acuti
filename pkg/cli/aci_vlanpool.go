package cli

import (
	"github.com/spf13/cobra"
)

// vpoolCmd represents the vlan-pool command
var vpoolCmd = &cobra.Command{
	Use:   "vlan-pool",
	Short: "Manage ACI vlan pools",
	Long:  `List, create & delete ACI vlan pools.`,
}

func init() {
	aciCmd.AddCommand(vpoolCmd)
}
