package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the ACI version number",
	Long:  `All software has versions. This is ACI's`,
	Run:   runVerCmd,
}

func init() {
	aciCmd.AddCommand(versionCmd)
}

func runVerCmd(cmd *cobra.Command, args []string) {
	fmt.Println("ACI Version here: ")
}
