package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// vrangeListCmd lists vlan ranges.
var vrangeListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists vlan ranges",
	Long:  `Lists vlan ranges`,
	Run:   runVrangeListCmd,
}

func init() {
	vrangeCmd.AddCommand(vrangeListCmd)
	vrangeListCmd.Flags().StringVar(&aciFlags.vpoolName, "vlan-pool", "", "Name of vlan pool")
	vrangeListCmd.Flags().StringVar(&aciFlags.vlanMode, "mode", "dynamic", "Allocation mode of the VLAN pool. Either dynamic or static.")
	vrangeListCmd.MarkFlagRequired("vlan-pool")
	vrangeListCmd.MarkFlagRequired("mode")
}

func runVrangeListCmd(cmd *cobra.Command, args []string) {
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	tw := newTabWriter(os.Stdout)
	defer tw.Flush()

	// Print the legend
	fmt.Fprintf(tw, "NAME\tDESCRIPTION\tSTATUS\tFROM\tTO\n")

	// ACI Client
	client := mustClientFromCmd(cmd)

	// Log into APIC API
	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Send vlan-range list request to the APIC API
	var resp []map[string]interface{}
	resp, err = client.VlanRangeList(aciFlags.vpoolName, aciFlags.vlanMode)
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", r["name"], r["descr"], r["status"], r["from"], r["to"])
	}
}
