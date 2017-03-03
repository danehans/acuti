package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vrangeDelCmd deletes an ACI vlan range.
var (
	vrangeDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Deletes an ACI vlan range.",
		Long:  `Deletes an ACI vlan range.`,
		Run:   RunVrangeDelCmd,
	}
)

func init() {
	vrangeCmd.AddCommand(vrangeDelCmd)
	vrangeDelCmd.Flags().StringVar(&aciFlags.vpoolName, "vlan-pool", "", "Name of vlan pool")
	vrangeDelCmd.Flags().StringVar(&aciFlags.vlanMode, "mode", "", "VLAN mode. TODO: Provide options.")
	vrangeDelCmd.Flags().StringVar(&aciFlags.vrangeStart, "start", "", "Starting VLAN number.")
	vrangeDelCmd.Flags().StringVar(&aciFlags.vrangeStop, "end", "", "Ending VLAN number.")
	vrangeDelCmd.MarkFlagRequired("vlan-pool")
	vrangeDelCmd.MarkFlagRequired("mode")
	vrangeDelCmd.MarkFlagRequired("start")
	vrangeDelCmd.MarkFlagRequired("end")
}

func RunVrangeDelCmd(cmd *cobra.Command, args []string) {
	if len(aciFlags.vpoolName) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.vlanMode) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.vrangeStart) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.vrangeStop) == 0 {
		cmd.Help()
		return
	}
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	client := mustClientFromCmd(cmd)

	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	// Note client package uses add naming instead of Del.
	err = client.VlanRangeDel(aciFlags.vpoolName, aciFlags.vlanMode, aciFlags.vrangeStart, aciFlags.vrangeStop)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("vlan range %s-%s deleted for vlan pool %s.\n", aciFlags.vrangeStart, aciFlags.vrangeStop, aciFlags.vpoolName)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
