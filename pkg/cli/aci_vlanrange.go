package cli

import (
	"github.com/spf13/cobra"
)

// vrangeCmd represents the vlan-range command
var vrangeCmd = &cobra.Command{
	Use:   "vlan-range",
	Short: "Manage ACI vlan ranges",
	Long:  `List, create & delete ACI vlan ranges.`,
}

func init() {
	aciCmd.AddCommand(vrangeCmd)
}
