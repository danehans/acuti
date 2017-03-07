package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// brCreateCmd creates a bridge domain.
var (
	brCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a bridge domain",
		Long:  `Create a bridge domain`,
		Run:   runBrCreate,
	}
)

func init() {
	brCmd.AddCommand(brCreateCmd)
	brCreateCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	brCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of bridge domain")
	brCreateCmd.Flags().StringVar(&aciFlags.descr, "description", "", "Description of bridge domain")
	brCreateCmd.MarkFlagRequired("tenant")
	brCreateCmd.MarkFlagRequired("name")
}

func runBrCreate(cmd *cobra.Command, args []string) {
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

	// Note client package uses add naming instead of create.
	err = client.BridgeDomainAdd(aciFlags.tenant, aciFlags.name, aciFlags.descr)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("bridge domain %s created for tenant %s.\n", aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
