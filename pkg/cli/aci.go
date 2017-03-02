package cli

import (
	"github.com/spf13/cobra"
)

// aciCmd represents the aci command
var (
	aciCmd = &cobra.Command{
		Use:   "aci",
		Short: "Manage aci...",
		Long:  `Do all sorts of stuff to aci`,
	}
	// aciFlags that can be set for any ACI subcommand.
	aciFlags = struct {
		name     string
		app      string
		tenant   string
		brdomain string
		descr    string
	}{}
)

func init() {
	RootCmd.AddCommand(aciCmd)
}
