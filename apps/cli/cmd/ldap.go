package cmd

import (
	"fmt"
	"os"

	"github.com/kubedev/kubeconfig-generator/pkg/client"
	"github.com/spf13/cobra"
)

var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "Generate the Kubernetes config for LDAP user token.",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient(flags)
		if err := c.GenerateKubeconfig(output); err != nil {
			fmt.Fprintf(os.Stderr, "%v.\n", err)
		}
	},
}

func init() {
	ldapCmd.Flags().StringVarP(&flags.DN, "dn", "", "", "Use the given DN to validate the target's LDAP server.")
	ldapCmd.Flags().StringVarP(&flags.Password, "password", "", "", "Use the given password to validate the target's LDAP server.")

	ldapCmd.MarkFlagRequired("dn")
	ldapCmd.MarkFlagRequired("password")
}
