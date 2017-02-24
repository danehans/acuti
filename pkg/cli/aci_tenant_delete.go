package cli

import (
	"github.com/spf13/cobra"
)

// tenantDeleteCmd creates an CI Tenant.
var (
	tenantDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Create an ACI tenant",
		Long:  `Create an ACI tenant`,
		Run:   runTenantDelete,
	}
	delName string
)

func init() {
	tenantCmd.AddCommand(tenantDeleteCmd)
	tenantDeleteCmd.Flags().StringVar(&delName, "name", "n", "Name of tenant")
	tenantDeleteCmd.MarkFlagRequired("name")
}

func runTenantDelete(cmd *cobra.Command, args []string) {
	if len(delName) == 0 {
		cmd.Help()
		return
	}
	if err := validateTenantDeleteArgs(cmd, args); err != nil {
		return
	}

	client := mustClientFromCmd(cmd)

	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Note client package uses add naming instead of create.
	err = client.TenantDel(delName)
	if err != nil {
		exitWithError(ExitError, err)
	}

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}

func validateTenantDeleteArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return usageError(cmd, "Unexpected args: %v", args)
	}
	return nil
}
