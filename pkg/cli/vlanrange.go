package cli

import (
	"github.com/spf13/cobra"
)

// vrangeCmd represents the vlan-range command
var vrangeCmd = &cobra.Command{
	Use:   "vlan-range",
	Short: "Manage vlan ranges",
	Long:  `List, create & delete vlan ranges.`,
}

func init() {
	RootCmd.AddCommand(vrangeCmd)
}
