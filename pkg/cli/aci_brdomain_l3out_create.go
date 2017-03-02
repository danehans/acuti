package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// brOutCreateCmd creates an ACI L3 outside network connection.
var (
	brOutCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an ACI bridge domain L3 outside network connection",
		Long:  `Create an ACI bridge domain L3 outside network connection`,
		Run:   RunBrOutCmd,
	}
)

func init() {
	brOutCmd.AddCommand(brOutCreateCmd)
	brOutCreateCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	brOutCreateCmd.Flags().StringVar(&aciFlags.brdomain, "bridge-domain", "", "Name of bridge domain")
	brOutCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of L3 outside network connection")
	brOutCreateCmd.MarkFlagRequired("tenant")
	brOutCreateCmd.MarkFlagRequired("bridge-domain")
	brOutCreateCmd.MarkFlagRequired("name")
}

// RunBrOutCmd creates an ACI L3 outside network connection
// for a tenant bridge domain.
func RunBrOutCmd(cmd *cobra.Command, args []string) {
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

	// Note client package uses add naming instead of create.
	err = client.BridgeDomainL3ExtOutAdd(aciFlags.tenant, aciFlags.brdomain, aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Created L3 outside network connection %s for tenant %s in bridge domain %s.\n",
		aciFlags.name, aciFlags.tenant, aciFlags.brdomain)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
