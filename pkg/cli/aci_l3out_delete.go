package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// l3outDelCmd deletes an ACI L3 outside network connection.
var (
	l3outDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an ACI L3 outside network connection",
		Long:  `Delete an ACI L3 outside network connection`,
		Run:   runL3outDel,
	}
)

func init() {
	l3outCmd.AddCommand(l3outDelCmd)
	l3outDelCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	l3outDelCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of L3 outside network connection")
	l3outDelCmd.MarkFlagRequired("tenant")
	l3outDelCmd.MarkFlagRequired("name")
}

func runL3outDel(cmd *cobra.Command, args []string) {
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

	err = client.L3ExtOutDel(aciFlags.tenant, aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Deleted L3 outside network connection %s for tenant %s.\n",
		aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
