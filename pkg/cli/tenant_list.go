package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// tenantListCmd lists tenants.
var tenantListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists tenants",
	Long:  `Lists tenants`,
	Run:   runTenantListCmd,
}

func init() {
	tenantCmd.AddCommand(tenantListCmd)
}

func runTenantListCmd(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}

	// Validate CLI arguments.
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	tw := newTabWriter(os.Stdout)
	defer tw.Flush()
	// legend
	fmt.Fprintf(tw, "NAME\tDESCRIPTION\n")

	// ACI Client
	client := mustClientFromCmd(cmd)

	//Log into ACI
	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Note client package uses add naming instead of create.
	var resp []map[string]interface{}
	resp, err = client.TenantList()
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["name"], r["descr"])
	}
}
