package ldap

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/ldap.v2"
)

type LdapUserInfo struct {
	KubernetesToken string
	Name            string
}

func dcToDomain(dc string) string {
  var buffer bytes.Buffer
  for index, domainString := range strings.Split(dc, ",") {
    if index != 0 {
      buffer.WriteString(".")
    }
    buffer.WriteString(strings.Split(domainString, "=")[1])
  }
  return buffer.String()
}

// QueryLdapUserInfo query user info from LDAP
func QueryLdapUserInfo(addr, dc, dn, passwd string) (*LdapUserInfo, error) {
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
	if err = l.Bind(fmt.Sprintf("%s@%s", dn, dcToDomain(dc)), passwd); err != nil {
		log.Println(err)
		return nil, err
	}

	// Search for the kubernetesToken and dn
	searchRequest := ldap.NewSearchRequest(
		dc,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(cn=%s))", dn),
		[]string{"cn", "objectGUID"},
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
		ldapUser.KubernetesToken = fmt.Sprintf("%x", entry.GetAttributeValue("objectGUID"))
		ldapUser.Name = entry.GetAttributeValue("cn")
	}
	return &ldapUser, nil
}
