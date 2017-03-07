package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tenantCreateCmd creates a tenant.
var (
	tenantCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a tenant",
		Long:  `Create a tenant`,
		Run:   runTenantCreate,
	}
)

func init() {
	tenantCmd.AddCommand(tenantCreateCmd)
	tenantCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of tenant")
	tenantCreateCmd.Flags().StringVar(&aciFlags.descr, "description", "", "Description of tenant")
	tenantCreateCmd.MarkFlagRequired("name")
}

func runTenantCreate(cmd *cobra.Command, args []string) {
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
	err = client.TenantAdd(aciFlags.name, aciFlags.descr)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Tenant %s created.\n", aciFlags.name)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
