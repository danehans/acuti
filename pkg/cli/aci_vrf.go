package cli

import (
	"github.com/spf13/cobra"
)

// vrfCmd represents the vrf command
var vrfCmd = &cobra.Command{
	Use:   "vrf",
	Short: "Manage ACI vrf's",
	Long:  `List, create & delete ACI vrf's.`,
}

func init() {
	aciCmd.AddCommand(vrfCmd)
}
