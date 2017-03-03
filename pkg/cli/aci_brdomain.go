package cli

import (
	"github.com/spf13/cobra"
)

// brCmd represents the bridge-domain command
var brCmd = &cobra.Command{
	Use:   "bridge-domain",
	Short: "Manage ACI bridge domains",
	Long:  `List, create & delete ACI bridge domains.`,
}

func init() {
	aciCmd.AddCommand(brCmd)
}
