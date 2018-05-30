package main

import (
	"fmt"
	"os"

	"github.com/inwinstack/kubeconfig-generator/apps/server/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kg",
	Short: "kg is the server for Kubeconfig Generator.",
}

func main() {
	addcommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addcommands() {
	rootCmd.AddCommand(cmd.ServeCmd)
	rootCmd.AddCommand(cmd.VersionCmd)
}
