package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// epgListCmd lists Endpoint Groups (EPGs).
var epgListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists Endpoint Groups (EPGs)",
	Long:  `Lists ACI Endpoint Groups (EPGs)`,
	Run:   runEPGListCmd,
}

func init() {
	epgCmd.AddCommand(epgListCmd)
	epgListCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	epgListCmd.Flags().StringVar(&aciFlags.app, "app", "", "Name of applicasstion profile")
	epgListCmd.MarkFlagRequired("tenant")
	epgListCmd.MarkFlagRequired("app")
}

func runEPGListCmd(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.app) == 0 {
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

	// Send epglication profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.ApplicationEPGList(aciFlags.tenant, aciFlags.app)
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["name"], r["descr"])
	}
}
