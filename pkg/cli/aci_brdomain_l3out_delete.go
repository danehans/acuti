package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// brOutDelCmd deletes an ACI bridge domain L3 outside network connection.
var (
	brOutDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an ACI bridge domain",
		Long:  `Delete an ACI bridge domain`,
		Run:   RunBrOutDelCmd,
	}
)

func init() {
	brOutCmd.AddCommand(brOutDelCmd)
	brOutDelCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	brOutDelCmd.Flags().StringVar(&aciFlags.brdomain, "bridge-domain", "", "Name of bridge domain")
	brOutDelCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of L3 outside network connection")
	brOutDelCmd.MarkFlagRequired("tenant")
	brOutDelCmd.MarkFlagRequired("bridge-domain")
	brOutDelCmd.MarkFlagRequired("name")
}

// RunBrOutDelCmd deletes an ACI L3 outside network connection
// for a tenant bridge domain.
func RunBrOutDelCmd(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.brdomain) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.name) == 0 {
		cmd.Help()
		return
	}
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	client := mustClientFromCmd(cmd)

	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	err = client.BridgeDomainL3ExtOutDel(aciFlags.tenant, aciFlags.brdomain, aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Deleted L3 outside network connection %s for tenant %s in bridge domain %s.\n",
		aciFlags.name, aciFlags.tenant, aciFlags.brdomain)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
