package cli

import (
	"github.com/spf13/cobra"
)

// epgCmd represents the epg command
var epgCmd = &cobra.Command{
	Use:   "epg",
	Short: "Manage ACI EPGs",
	Long:  `List, create & delete ACI EPGs.`,
}

func init() {
	aciCmd.AddCommand(epgCmd)
}
