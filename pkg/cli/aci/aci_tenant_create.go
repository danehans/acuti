package cli

import (
	"github.com/spf13/cobra"
)

// tenantAddCmd creates an CI Tenant.
var (
	tenantAddCmd = &cobra.Command{
		Use:   "add name description",
		Short: "Add an ACI tenant",
		Long:  `Add an ACI tenant`,
		Run:   runTenantAdd,
	}
	name        string
	description string
)

func init() {
	tenantCmd.AddCommand(tenantAddCmd)
	tenantAddCmd.Flags().StringVar(&name, "name", "n", "Name of tenant")
	tenantAddCmd.Flags().StringVar(&description, "description", "d", "Description of tenant")
	tenantAddCmd.MarkFlagRequired("name")
}

func runTenantAdd(cmd *cobra.Command, args []string) {
	if len(name) == 0 {
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

	err = client.TenantAdd(name, description)
	if err != nil {
		exitWithError(ExitError, err)
	}

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return usageError(cmd, "Unexpected args: %v", args)
	}
	return nil
}
