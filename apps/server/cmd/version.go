package cmd

import (
	"fmt"

	"github.com/inwinstack/kubeconfig-generator/pkg/version"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s\n", version.GetVersion())
		return nil
	},
}
