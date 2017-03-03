package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vpoolCreateCmd creates an ACI vlan pool.
var (
	vpoolCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an ACI vlan pool",
		Long:  `Create an ACI vlan pool`,
		Run:   RunVpoolCreateCmd,
	}
)

func init() {
	vpoolCmd.AddCommand(vpoolCreateCmd)
	vpoolCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of vlan pool")
	vpoolCreateCmd.Flags().StringVar(&aciFlags.vlanMode, "mode", "dynamic", "Allocation mode of the VLAN pool. Either dynamic or static.")
	vpoolCreateCmd.Flags().StringVar(&aciFlags.descr, "description", "", "Description of vlan pool")
	vpoolCreateCmd.MarkFlagRequired("name")
	vpoolCreateCmd.MarkFlagRequired("mode")
}

func RunVpoolCreateCmd(cmd *cobra.Command, args []string) {
	if len(aciFlags.name) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.vlanMode) == 0 {
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
	err = client.VlanPoolAdd(aciFlags.name, aciFlags.vlanMode, aciFlags.descr)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("vlan pool %s created.\n", aciFlags.name)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
