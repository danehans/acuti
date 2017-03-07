package cli

import (
	"github.com/spf13/cobra"
)

// vrfCmd represents the virtual routing and forwarding (vrf) command
var vrfCmd = &cobra.Command{
	Use:   "vrf",
	Short: "Manage virtual routing and forwarding (VRFs)",
	Long:  `List, create & delete virtual routing and forwarding (VRF's).`,
}

func init() {
	RootCmd.AddCommand(vrfCmd)
}
