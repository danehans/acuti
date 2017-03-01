package cli

import (
	"github.com/spf13/cobra"
)

// appCmd represents the application profile command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Manage ACI application profiles",
	Long:  `List, create & delete ACI application profiles.`,
}

func init() {
	aciCmd.AddCommand(appCmd)
}
