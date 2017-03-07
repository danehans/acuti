package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/danehans/acigo/aci"
)

var (
	// RootCmd is the base acictl command.
	RootCmd = &cobra.Command{
		Use:   "acictl",
		Short: "A command line utility for managing Cisco ACI",
		Long: `A command line utility for managing Cisco ACI

To get help about a resource or command, run "acictl help resource"`,
	}

	// globalFlags can be set for any subcommand.
	globalFlags = struct {
		hosts    []string
		username string
		password string
	}{}

	// aciFlags that can be set for any ACI subcommand.
	aciFlags = struct {
		name        string
		app         string
		tenant      string
		brdomain    string
		vpoolName   string
		vlanMode    string
		vrangeStart string
		vrangeStop  string
		descr       string
	}{}
)

func init() {
	RootCmd.PersistentFlags().StringSliceVar(&globalFlags.hosts, "hosts", []string{"sandboxapicdc.cisco.com"}, "APIC API Endpoints")
	// gRPC TLS Server Verification
	RootCmd.PersistentFlags().StringVar(&globalFlags.username, "user", "admin", "APIC Username")
	// gRPC TLS Client Authentication
	RootCmd.PersistentFlags().StringVar(&globalFlags.password, "pass", "C1sco12345", "APIC Password")
	cobra.EnablePrefixMatching = true
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// mustClientFromCmd returns an ACI client or exits.
func mustClientFromCmd(cmd *cobra.Command) *aci.Client {
	hosts := hostsFromCmd(cmd)
	user := userFromCmd(cmd)
	pass := passFromCmd(cmd)

	// Sets ACI client configuration options.
	opts := aci.ClientOptions{
		Hosts: hosts,
		User:  user,
		Pass:  pass,
	}

	// Creates an instance of the ACI Client.
	client, err := aci.New(opts)
	if err != nil {
		exitWithError(ExitBadConnection, err)
	}
	return client
}

// hostsFromCmd returns the hosts argument.
func hostsFromCmd(cmd *cobra.Command) []string {
	hosts, err := cmd.Flags().GetStringSlice("hosts")
	if err != nil {
		exitWithError(ExitBadArgs, err)
	}
	return hosts
}

// userFromCmd returns the user argument.
func userFromCmd(cmd *cobra.Command) string {
	user, err := cmd.Flags().GetString("user")
	if err != nil {
		exitWithError(ExitBadArgs, err)
	}
	return user
}

// passFromCmd returns the password argument.
func passFromCmd(cmd *cobra.Command) string {
	password, err := cmd.Flags().GetString("pass")
	if err != nil {
		exitWithError(ExitBadArgs, err)
	}
	return password
}

// TODO: Remove or refactor arg validation.
func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return usageError(cmd, "Unexpected args: %v", args)
	}
	return nil
}
