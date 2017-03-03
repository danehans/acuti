package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// nodeListCmd describes a Node.
var nodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACI nodes",
	Long:  `Lists ACI nodes`,
	Run:   RunNodeListCmd,
}

func init() {
	nodeCmd.AddCommand(nodeListCmd)
}

func RunNodeListCmd(cmd *cobra.Command, args []string) {
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

	// Send node profile list request to APIC API
	var resp []map[string]interface{}
	resp, err = client.NodeList()
	if err != nil {
		exitWithError(ExitError, err)
	}
	for _, r := range resp {
		fmt.Fprintf(tw, "%s\t%s\n", r["name"], r["descr"])
	}
}
