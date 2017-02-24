package aci

import (
	"github.com/spf13/cobra"

	"github.com/danehans/octogo/pkg/cli"
)

// AciCmd represents the aci command
var AciCmd = &cobra.Command{
	Use:   "aci",
	Short: "Manage aci...",
	Long:  `Do all sorts of stuff to aci`,
}

func init() {
	cli.RootCmd.AddCommand(AciCmd)
}
