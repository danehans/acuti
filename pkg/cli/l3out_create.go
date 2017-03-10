package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// l3outCreateCmd creates a L3 outside network connection.
var (
	l3outCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a L3 outside network connection",
		Long:  `Create a L3 outside network connection`,
		Run:   runL3outCreateCmd,
	}
)

func init() {
	l3outCmd.AddCommand(l3outCreateCmd)
	l3outCreateCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	l3outCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of L3 outside network connection")
	l3outCreateCmd.Flags().StringVar(&aciFlags.descr, "descr", "", "Description of L3 outside network connection")
	l3outCreateCmd.MarkFlagRequired("tenant")
	l3outCreateCmd.MarkFlagRequired("name")
}

func runL3outCreateCmd(cmd *cobra.Command, args []string) {
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
	err = client.L3ExtOutAdd(aciFlags.tenant, aciFlags.name, aciFlags.descr)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Created L3 outside network connection %s for tenant %s.\n",
		aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
