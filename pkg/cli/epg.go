package cli

import (
	"github.com/spf13/cobra"
)

// epgCmd represents the epg command
var epgCmd = &cobra.Command{
	Use:   "epg",
	Short: "Manage EPGs",
	Long:  `List, create & delete EPGs.`,
}

func init() {
	RootCmd.AddCommand(epgCmd)
}
