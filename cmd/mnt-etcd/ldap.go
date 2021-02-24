package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/freman/caddy2-reauth/backends/ldap"
	"github.com/freman/caddy2-reauth/jsontypes"
)

func ldapMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		driver := getDriver()
		userDN, err := driver.Authenticate(r)
		if userDN == "" || err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
			http.Error(w, "credential invalid", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ldapHandler(driver *ldap.LDAP) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userDN, err := driver.Authenticate(r)
		if userDN == "" || err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
			http.Error(w, "credential invalid", http.StatusUnauthorized)
			return
		}
		io.WriteString(w, fmt.Sprintf("%s\n", userDN))
	})
}

func getDriver() *ldap.LDAP {
	u, err := url.Parse("ldap://ldap.pow.ai:389")
	if err != nil {
		panic(err)
	}

	driver := ldap.NewDriver()
	driver.URL = &jsontypes.URL{u}
	driver.BindDN = "cn=admin,dc=pow,dc=ai"
	driver.BindPassword = "pow4506"
	driver.FilterDN = "(&(objectClass=inetOrgPerson)(cn=%s))"
	driver.BaseDN = "ou=People,dc=pow,dc=ai"

	if err := driver.Validate(); err != nil {
		panic(err)
	}
	return driver
}
