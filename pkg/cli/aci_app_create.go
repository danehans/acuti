package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// appCreateCmd creates an ACI Application Profile.
var (
	appCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an ACI Application Profile",
		Long:  `Create an ACI Application Profile`,
		Run:   runAppCreate,
	}
)

func init() {
	appCmd.AddCommand(appCreateCmd)
	appCreateCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	appCreateCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of application profile")
	appCreateCmd.Flags().StringVar(&aciFlags.descr, "description", "", "Description of app profile")
	appCreateCmd.MarkFlagRequired("tenant")
	appCreateCmd.MarkFlagRequired("name")
}

func runAppCreate(cmd *cobra.Command, args []string) {
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
	err = client.ApplicationProfileAdd(aciFlags.tenant, aciFlags.name, aciFlags.descr)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Application profile %s created for tenant %s.\n", aciFlags.name, aciFlags.tenant)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
