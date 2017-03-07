package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// brListCmd lists bridge domains for a tenant.
var brListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists bridge domains for a tenant",
	Long:  `Lists bridge domains for a tenant`,
	Run:   runBrListCmd,
}

func init() {
	brCmd.AddCommand(brListCmd)
	brListCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	brListCmd.MarkFlagRequired("tenant")
}

func runBrListCmd(cmd *cobra.Command, args []string) {
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

	// Send bridge profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.BridgeDomainList(aciFlags.tenant)
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["name"], r["descr"])
	}
}
