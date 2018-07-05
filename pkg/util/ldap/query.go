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
func QueryLdapUserInfo(addr, dc, userSearchBase, userNameAttribute, userTokenAttribute, dn, passwd string) (*LdapUserInfo, error) {
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
	ldap_user := dn
	user_search_base := dn
	user_search_filter := "(&(objectClass=*))"
  if !strings.Contains(strings.ToLower(dn), strings.ToLower(dc)) {
		ldap_user = fmt.Sprintf("%s@%s", dn, dcToDomain(dc))
		user_search_filter = fmt.Sprintf("(&(userPrincipalName=%s))", ldap_user)
		user_search_base = userSearchBase
		if len(userSearchBase) == 0 {
			user_search_base = dc
		}
	}

	if err = l.Bind(ldap_user, passwd); err != nil {
		log.Println(err)
		return nil, err
	}

	// Search for the kubernetesToken and dn
	searchRequest := ldap.NewSearchRequest(
		user_search_base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		user_search_filter,
		[]string{userNameAttribute, userTokenAttribute},
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
		ldapUser.KubernetesToken = entry.GetAttributeValue(userTokenAttribute)
		if userTokenAttribute == "objectGUID" {
			ldapUser.KubernetesToken = fmt.Sprintf("%x", entry.GetAttributeValue(userTokenAttribute))
		}
		ldapUser.Name = entry.GetAttributeValue(userNameAttribute)
	}
	return &ldapUser, nil
}
