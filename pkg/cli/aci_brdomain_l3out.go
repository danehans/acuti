package cli

import (
	"github.com/spf13/cobra"
)

// BrOutCmd represents the bridge-domain l3-out command
var brOutCmd = &cobra.Command{
	Use:   "l3-out",
	Short: "Manage ACI tenant outside bridge domain network connections",
	Long:  `List, create & delete ACI tenant layer-3 outside network connections for a bridge domain.`,
}

func init() {
	brCmd.AddCommand(brOutCmd)
}
