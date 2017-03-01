package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// appListCmd describes a Group.
var appListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACI application profiles for a tenant",
	Long:  `Lists ACI application profiles for a tenant`,
	Run:   runAppListCmd,
}

func init() {
	appCmd.AddCommand(appListCmd)
	appListCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "t", "Name of tenant")
	appListCmd.MarkFlagRequired("tenant")
}

func runAppListCmd(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
		cmd.Help()
		return
	}

	// Validate CLI arguments.
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	tw := newTabWriter(os.Stdout)
	defer tw.Flush()

	// Print the legend
	fmt.Fprintf(tw, "NAME\tDESCRIPTION\n")

	// ACI Client
	client := mustClientFromCmd(cmd)

	//Log into APIC API
	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Send application profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.ApplicationProfileList(aciFlags.tenant)
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["name"], r["descr"])
	}
}
