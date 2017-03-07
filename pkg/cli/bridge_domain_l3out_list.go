package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// brOutListCmd lists ACI L3 outside network connections
// for a tenant's bridge domain.
var brOutListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACI bridge domain L3 outside network connections",
	Long:  `Lists ACI bridge L3 outside network connections for a tenant's bridge domain.`,
	Run:   runBrOutListCmd,
}

func init() {
	brOutCmd.AddCommand(brOutListCmd)
	brOutListCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	brOutListCmd.Flags().StringVar(&aciFlags.brdomain, "bridge-domain", "", "Name of bridge-domain")
	brOutListCmd.MarkFlagRequired("tenant")
	brOutListCmd.MarkFlagRequired("bridge-domain")
}

func runBrOutListCmd(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.brdomain) == 0 {
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
	fmt.Fprintf(tw, "NAME\tSTATE\n")

	// ACI Client
	client := mustClientFromCmd(cmd)

	//Log into APIC API
	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Send vrflication profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.BridgeDomainL3ExtOutList(aciFlags.tenant, aciFlags.brdomain)
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["tnL3extOutName"], r["state"])
	}
}
