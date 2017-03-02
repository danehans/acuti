package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// l3ListCmd lists ACI L3 outside network connections for a tenant.
var l3ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACI L3 outside network connections for a tenant",
	Long:  `Lists ACI L3 outside network connections for a tenant.`,
	Run:   runL3outListCmd,
}

func init() {
	l3outCmd.AddCommand(l3ListCmd)
	l3ListCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	l3ListCmd.MarkFlagRequired("tenant")
}

func runL3outListCmd(cmd *cobra.Command, args []string) {
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
	fmt.Fprintf(tw, "NAME\tSTATUS\tDESCRIPTION\n")

	// ACI Client
	client := mustClientFromCmd(cmd)

	//Log into APIC API
	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Send vrflication profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.L3ExtOutList(aciFlags.tenant)
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", r["name"], r["status"], r["descr"])
	}
}
