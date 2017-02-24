package aci

import (
	"github.com/spf13/cobra"

	"github.com/danehans/octogo/pkg/cli"
)

// tenantCmd represents the tenant command
var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Manage ACI tenants",
	Long:  `List, create & delete ACI tenants`,
}

func init() {
	cli.aciCmd.AddCommand(tenantCmd)
}
