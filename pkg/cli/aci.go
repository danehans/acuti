package cli

import (
	"github.com/spf13/cobra"
)

// aciCmd represents the aci command
var aciCmd = &cobra.Command{
	Use:   "aci",
	Short: "Manage aci...",
	Long:  `Do all sorts of stuff to aci`,
}

func init() {
	RootCmd.AddCommand(aciCmd)
}
