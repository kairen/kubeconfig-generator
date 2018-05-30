package main

import (
	"fmt"

	"os"

	"github.com/inwinstack/kubeconfig-generator/apps/cli/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kgctl",
	Short: "kgctl is the command line tool for Kubeconfig Generator.",
}

func main() {
	addcommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addcommands() {
	rootCmd.AddCommand(cmd.GenerateCmd)
	rootCmd.AddCommand(cmd.VersionCmd)
}
