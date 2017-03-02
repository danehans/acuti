package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// vrfListCmd lists vrf's for a tenant.
var vrfListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACI vrf's for a tenant",
	Long:  `Lists ACI vrf's for a tenant`,
	Run:   runVRFListCmd,
}

func init() {
	vrfCmd.AddCommand(vrfListCmd)
	vrfListCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	vrfListCmd.MarkFlagRequired("tenant")
}

func runVRFListCmd(cmd *cobra.Command, args []string) {
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

	// Send vrf profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.VrfList(aciFlags.tenant)
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["name"], r["descr"])
	}
}
