package cmd

import (
	"fmt"
	"os"

	"github.com/inwinstack/kubeconfig-generator/pkg/server"
	"github.com/spf13/cobra"
)

var (
	listen             string
	endpoint           string
	caPath             string
	ldapAddr           string
	ldapDC	           string
	userSearchBase     string
	userNameAttribute  string
	userTokenAttribute string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the generator server.",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(listen, endpoint, caPath, ldapAddr, ldapDC, userSearchBase, userNameAttribute, userTokenAttribute)
		if err := s.Serve(); err != nil {
			fmt.Fprintf(os.Stderr, "Error serving: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	serveCmd.Flags().StringVarP(&listen, "listen", "", ":8080", "Server serve address.")
	serveCmd.Flags().StringVarP(&endpoint, "kube-apiserver-endpoint", "", "", "Kubernetes API server external endpoint.")
	serveCmd.Flags().StringVarP(&caPath, "ca-path", "", "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt", "Kubernetes API server certificate.")
	serveCmd.Flags().StringVarP(&ldapAddr, "ldap-address", "", "", "LDAP server address.")
	serveCmd.Flags().StringVarP(&ldapDC, "ldap-dc", "", "", "LDAP domain componet(DC).")
	serveCmd.Flags().StringVarP(&userSearchBase, "user-search-base", "", "", "Base distinguished name(DN) to search for users in.")
	serveCmd.Flags().StringVarP(&userNameAttribute, "user-name-attribute", "", "", "Attribute of the user entry that contains their username.")
	serveCmd.Flags().StringVarP(&userTokenAttribute, "user-token-arttribute", "", "", "Attribute of the user entry that contains their token.")
	serveCmd.MarkFlagRequired("kube-apiserver-endpoint")
	serveCmd.MarkFlagRequired("ldap-dc")
	serveCmd.MarkFlagRequired("user-name-attribute")
	serveCmd.MarkFlagRequired("user-token-arttribute")
}
