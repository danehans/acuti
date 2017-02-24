package cli

import (
	"github.com/spf13/cobra"
)

// tenantCmd represents the tenant command
var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Manage ACI tenants",
	Long:  `List, create & delete ACI tenants`,
}

func init() {
	aciCmd.AddCommand(tenantCmd)
}
