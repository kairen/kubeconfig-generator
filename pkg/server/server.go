package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inwinstack/kubeconfig-generator/pkg/types"
	"github.com/inwinstack/kubeconfig-generator/pkg/util"
	"github.com/inwinstack/kubeconfig-generator/pkg/util/ldap"
)

type Server struct {
	router   *gin.Engine
	listen   string
	apiURL   string
	caPath   string
	ldapAddr string
}

func NewServer(listen, apiURL, caPath, ldapAddr string) *Server {
	gin.DisableConsoleColor()
	server := &Server{
		router:   gin.Default(),
		listen:   listen,
		apiURL:   apiURL,
		caPath:   caPath,
		ldapAddr: ldapAddr,
	}
	return server
}

func (s *Server) Serve() error {

	// healthz check
	s.router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// login to query token
	s.router.POST("/login", func(c *gin.Context) {
		var user types.User
		if err := c.ShouldBindJSON(&user); err == nil {
			ldapUser, qerr := ldap.QueryLdapUserInfo(s.ldapAddr, user.DN, user.Password)
			ca, caerr := util.LoadBase64CertificateAuthority(s.caPath)
			if ldapUser != nil && qerr == nil && caerr == nil {
				c.JSON(http.StatusOK, types.Generator{
					UserName: ldapUser.Name,
					Token:    ldapUser.KubernetesToken,
					CA:       ca,
					Endpoint: s.apiURL,
					Status:   types.Authorized,
				})
			} else {
				c.JSON(http.StatusUnauthorized, types.Generator{Status: types.Unauthorized})
			}
		} else {
			c.JSON(http.StatusBadRequest, types.Generator{
				Status:  types.Error,
				Message: err.Error(),
			})
		}
	})

	if err := s.router.Run(s.listen); err != nil {
		return err
	}
	return nil
}
