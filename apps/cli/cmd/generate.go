package cmd

import (
	"fmt"
	"os"

	"github.com/inwinstack/kubeconfig-generator/pkg/client"
	"github.com/spf13/cobra"
)

var (
	url      string
	password string
	dn       string
	output   string
)

var GenerateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate the Kubernetes config from the server.",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient(url, dn, password)
		if err := c.GenerateKubeconfig(output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&url, "url", "", "", "Kubeconfig generator server url.")
	GenerateCmd.Flags().StringVarP(&output, "output", "o", "~/.kube/config", "Kubeconfig output path.")
	GenerateCmd.Flags().StringVarP(&dn, "dn", "", "", "LDAP server login domain name.")
	GenerateCmd.Flags().StringVarP(&password, "password", "", "", "LDAP server login password.")
	GenerateCmd.MarkFlagRequired("url")
	GenerateCmd.MarkFlagRequired("dn")
	GenerateCmd.MarkFlagRequired("password")
}
