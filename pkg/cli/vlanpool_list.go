package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// vpoolListCmd lists vlan pools.
var vpoolListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists vlan pools",
	Long:  `Lists vlan pools`,
	Run:   runVpoolListCmd,
}

func init() {
	vpoolCmd.AddCommand(vpoolListCmd)
}

func runVpoolListCmd(cmd *cobra.Command, args []string) {
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

	// Send vpoollication profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.VlanPoolList()
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["name"], r["descr"])
	}
}
