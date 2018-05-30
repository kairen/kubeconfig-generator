package cmd

import (
	"fmt"
	"os"

	"github.com/inwinstack/kubeconfig-generator/pkg/server"
	"github.com/spf13/cobra"
)

var (
	listen   string
	endpoint string
	caPath   string
	ldapAddr string
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the generator server.",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(listen, endpoint, caPath, ldapAddr)
		if err := s.Serve(); err != nil {
			fmt.Fprintf(os.Stderr, "Error serving: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	ServeCmd.Flags().StringVarP(&listen, "listen", "", ":8080", "Server serve address.")
	ServeCmd.Flags().StringVarP(&endpoint, "kube-apiserver-endpoint", "", "", "Kubernetes API server external endpoint.")
	ServeCmd.Flags().StringVarP(&caPath, "ca-path", "", "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt", "Kubernetes API server certificate.")
	ServeCmd.Flags().StringVarP(&ldapAddr, "ldap-address", "", "", "LDAP server address.")
	ServeCmd.MarkFlagRequired("kube-apiserver-endpoint")
	ServeCmd.MarkFlagRequired("ldap-address")
}
