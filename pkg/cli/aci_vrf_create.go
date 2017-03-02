package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vrfCreateCmd creates an ACI vrf.
var (
	vrfCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an ACI vrf",
		Long:  `Create an ACI vrf`,
		Run:   runVRFCreate,
	}
)

func init() {
	vrfCmd.AddCommand(vrfCreateCmd)
	vrfCreateCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	vrfCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of vrf")
	vrfCreateCmd.Flags().StringVar(&aciFlags.descr, "description", "", "Description of vrf")
	vrfCreateCmd.MarkFlagRequired("tenant")
	vrfCreateCmd.MarkFlagRequired("name")
}

func runVRFCreate(cmd *cobra.Command, args []string) {
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

	// Note client package uses add naming instead of create.
	err = client.VrfAdd(aciFlags.tenant, aciFlags.name, aciFlags.descr)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("vrf %s created for tenant %s.\n", aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
