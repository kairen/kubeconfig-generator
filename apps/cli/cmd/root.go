package cmd

import (
	"fmt"
	"os"

	"github.com/kairen/kubeconfig-generator/pkg/client"
	"github.com/spf13/cobra"
)

var (
	flags  client.Flags
	output string
)

var rootCmd = &cobra.Command{
	Use:   "kgctl",
	Short: "kgctl is the command line tool for Kubeconfig Generator.",
}

func Execute() {
	addCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addCommands() {
	rootCmd.AddCommand(ldapCmd)
	rootCmd.AddCommand(versionCmd)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "~/.kube/config", "Kubeconfig output path.")
	rootCmd.PersistentFlags().StringVarP(&flags.URL, "url", "", "http://localhost:8080", "A server endpoint for Kubeconfig Generator.")
}
