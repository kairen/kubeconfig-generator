package ldap

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/ldap.v2"
)

type LdapUserInfo struct {
	KubernetesToken string
	Name            string
}

// QueryLdapUserInfo query user info from LDAP
func QueryLdapUserInfo(addr, dn, passwd string) (*LdapUserInfo, error) {
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LDAP-debug] ")

	l, err := ldap.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer l.Close()

	if ok, _ := strconv.ParseBool(os.Getenv("ENABLE_START_TLS")); ok {
		if err = l.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
			return nil, err
		}
	}

	// login LDAP server by dn and password
	if err = l.Bind(dn, passwd); err != nil {
		log.Println(err)
		return nil, err
	}

	// Search for the kubernetesToken and dn
	searchRequest := ldap.NewSearchRequest(
		"dc=k8s,dc=com", // need change
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(sAMAccountName=%s))", dn),
		[]string{"sAMAccountName", "whenCreated"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(sr.Entries) != 1 {
		err := fmt.Errorf("User does not exist or too many entries returned")
		log.Println(err)
		return nil, err
	}

	var ldapUser LdapUserInfo
	for _, entry := range sr.Entries {
		name := entry.GetAttributeValue("sAMAccountName")
		whenCreated := entry.GetAttributeValue("whenCreated")
		token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s-%s", name, whenCreated)))
		ldapUser.KubernetesToken = token
		ldapUser.Name = name
	}
	return &ldapUser, nil
}
