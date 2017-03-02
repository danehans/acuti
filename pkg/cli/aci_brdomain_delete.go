package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// brDelCmd deletes an ACI bridge domain.
var (
	brDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an ACI bridge domain",
		Long:  `Delete an ACI bridge domain`,
		Run:   runBrDel,
	}
)

func init() {
	brCmd.AddCommand(brDelCmd)
	brDelCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	brDelCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of bridge domain")
	brDelCmd.MarkFlagRequired("tenant")
	brDelCmd.MarkFlagRequired("name")
}

func runBrDel(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
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

	err = client.BridgeDomainDel(aciFlags.tenant, aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Bridge domain %s deleted for tenant %s.\n", aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
