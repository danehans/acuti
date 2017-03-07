package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// epgDelCmd deletes an Endpoint Group (EPG).
var (
	epgDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an Endpoint Group (EPG)",
		Long:  `Delete an Endpoint Group (EPG)`,
		Run:   runEPGDel,
	}
)

func init() {
	epgCmd.AddCommand(epgDelCmd)
	epgDelCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	epgDelCmd.Flags().StringVar(&aciFlags.app, "app", "", "Name of application profile")
	epgDelCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of EPG")
	epgDelCmd.MarkFlagRequired("tenant")
	epgDelCmd.MarkFlagRequired("name")
}

func runEPGDel(cmd *cobra.Command, args []string) {
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

	err = client.ApplicationEPGDel(aciFlags.tenant, aciFlags.app, aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("EPG %s deleted for tenant %s.\n", aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
