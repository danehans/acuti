package cli

import (
	"github.com/spf13/cobra"
)

// tenantCreateCmd creates an CI Tenant.
var (
	tenantCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an ACI tenant",
		Long:  `Create an ACI tenant`,
		Run:   runTenantCreate,
	}
	addName     string
	description string
)

func init() {
	tenantCmd.AddCommand(tenantCreateCmd)
	tenantCreateCmd.Flags().StringVar(&addName, "name", "n", "Name of tenant")
	tenantCreateCmd.Flags().StringVar(&description, "description", "d", "Description of tenant")
	tenantCreateCmd.MarkFlagRequired("name")
}

func runTenantCreate(cmd *cobra.Command, args []string) {
	if len(addName) == 0 {
		cmd.Help()
		return
	}
	if err := validateTenantCreateArgs(cmd, args); err != nil {
		return
	}

	client := mustClientFromCmd(cmd)

	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Note client package uses add naming instead of create.
	err = client.TenantAdd(addName, description)
	if err != nil {
		exitWithError(ExitError, err)
	}

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}

func validateTenantCreateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return usageError(cmd, "Unexpected args: %v", args)
	}
	return nil
}
