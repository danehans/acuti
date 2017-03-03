package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// epgCreateCmd creates an ACI EPG.
var (
	epgCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an ACI epg",
		Long:  `Create an ACI epg`,
		Run:   RunEPGCreate,
	}
)

func init() {
	epgCmd.AddCommand(epgCreateCmd)
	epgCreateCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	epgCreateCmd.Flags().StringVar(&aciFlags.app, "app", "", "Name of application profile")
	epgCreateCmd.Flags().StringVar(&aciFlags.brdomain, "bridge-domain", "", "Name of bridge domain")
	epgCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of epg")
	epgCreateCmd.Flags().StringVar(&aciFlags.descr, "description", "", "Description of epg")
	epgCreateCmd.MarkFlagRequired("tenant")
	epgCreateCmd.MarkFlagRequired("name")
	epgCreateCmd.MarkFlagRequired("bridge")
}

func RunEPGCreate(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.name) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.brdomain) == 0 {
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
	err = client.ApplicationEPGAdd(aciFlags.tenant, aciFlags.app, aciFlags.brdomain, aciFlags.name, aciFlags.descr)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("epg %s created.\n", aciFlags.name)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
