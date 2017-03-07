package cli

import (
	"github.com/spf13/cobra"
)

// BrOutCmd represents the bridge-domain l3-out command
var brOutCmd = &cobra.Command{
	Use:   "l3-out",
	Short: "Manage tenant layer-3 outside network connections for a bridge domain",
	Long:  `List, create & delete tenant layer-3 outside network connections for a bridge domain.`,
}

func init() {
	brCmd.AddCommand(brOutCmd)
}
