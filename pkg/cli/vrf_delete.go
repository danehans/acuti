package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vrfDelCmd deletes a virtual routing and forwarding (vrf) instance.
var (
	vrfDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a virtual routing and forwarding (vrf) instance",
		Long:  `Delete a virtual routing and forwarding (vrf) instance`,
		Run:   runVRFDel,
	}
)

func init() {
	vrfCmd.AddCommand(vrfDelCmd)
	vrfDelCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	vrfDelCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of vrf")
	vrfDelCmd.MarkFlagRequired("tenant")
	vrfDelCmd.MarkFlagRequired("name")
}

func runVRFDel(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.name) == 0 {
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

	err = client.VrfDel(aciFlags.tenant, aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("vrf %s deleted for tenant %s.\n", aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
