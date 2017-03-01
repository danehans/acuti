package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tenantDeleteCmd creates an CI Tenant.
var (
	tenantDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an ACI tenant",
		Long:  `Delete an ACI tenant`,
		Run:   runTenantDelete,
	}
)

func init() {
	tenantCmd.AddCommand(tenantDeleteCmd)
	tenantDeleteCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of tenant")
	tenantDeleteCmd.MarkFlagRequired("name")
}

func runTenantDelete(cmd *cobra.Command, args []string) {
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
	err = client.TenantDel(aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Tenant %s deleted.\n", aciFlags.name)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
