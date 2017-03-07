package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vrangeCreateCmd creates a vlan range.
var (
	vrangeCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a vlan range",
		Long:  `Create a vlan range`,
		Run:   runVrangeCreateCmd,
	}
)

func init() {
	vrangeCmd.AddCommand(vrangeCreateCmd)
	vrangeCreateCmd.Flags().StringVar(&aciFlags.vpoolName, "vlan-pool", "", "Name of vlan pool")
	vrangeCreateCmd.Flags().StringVar(&aciFlags.vlanMode, "mode", "", "VLAN mode.")
	vrangeCreateCmd.Flags().StringVar(&aciFlags.vrangeStart, "start", "", "Starting VLAN number.")
	vrangeCreateCmd.Flags().StringVar(&aciFlags.vrangeStop, "end", "", "Ending VLAN number.")
	vrangeCreateCmd.MarkFlagRequired("vlan-pool")
	vrangeCreateCmd.MarkFlagRequired("mode")
	vrangeCreateCmd.MarkFlagRequired("start")
	vrangeCreateCmd.MarkFlagRequired("end")
}

func runVrangeCreateCmd(cmd *cobra.Command, args []string) {
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

	// Note client package uses add naming instead of create.
	err = client.VlanRangeAdd(aciFlags.vpoolName, aciFlags.vlanMode, aciFlags.vrangeStart, aciFlags.vrangeStop)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("vlan range %s-%s created for vlan pool %s.\n", aciFlags.vrangeStart, aciFlags.vrangeStop, aciFlags.vpoolName)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
