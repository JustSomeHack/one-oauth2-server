package ldapconfig

import (
	"log"
	"os"
	"strconv"
)

// LDAPConfig configurations for LDAP server connection
type LDAPConfig struct {
	BaseDN       string
	BindUser     string
	BindPassword string
	LDAPServer   string
	LDAPPort     int
	UseTLS       bool
	StartTLS     bool
}

// LoadLDAPConfig loads configuration from environment variable
func LoadLDAPConfig() *LDAPConfig {
	port, err := strconv.Atoi(os.Getenv("LDAP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	useTLS, err := strconv.ParseBool(os.Getenv("LDAP_USE_TLS"))
	if err != nil {
		useTLS = false
	}

	return &LDAPConfig{
		BaseDN:       os.Getenv("LDAP_BASE_DN"),
		BindUser:     os.Getenv("LDAP_BIND_USER"),
		BindPassword: os.Getenv("LDAP_BIND_PASSWORD"),
		LDAPServer:   os.Getenv("LDAP_SERVER"),
		LDAPPort:     port,
		UseTLS:       useTLS,
	}
}
