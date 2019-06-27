package cmd

import (
	"fmt"

	"github.com/kairen/kubeconfig-generator/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of client.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s\n", version.GetVersion())
		return nil
	},
}
