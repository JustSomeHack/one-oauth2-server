package ldapservice

import (
	"crypto/tls"
	"fmt"

	"github.com/JustSomeHack/one-oauth2-server/models"
	"github.com/JustSomeHack/one-oauth2-server/models/ldapconfig"

	"gopkg.in/ldap.v3"
)

// LDAPService service to connect to LDAP server
type LDAPService interface {
	Authenticate(username string, password string) (*models.User, error)
	bind() (err error)
}

type ldapService struct {
	baseDN       string
	bindUser     string
	bindPassword string
	ldapServer   string
	ldapPort     int
	useTLS       bool
	startTLS     bool
	conn         *ldap.Conn
}

// NewLDAPService gets a new reference to LDAPService
func NewLDAPService(ldapConfig *ldapconfig.LDAPConfig) LDAPService {
	return &ldapService{
		baseDN:       ldapConfig.BaseDN,
		bindUser:     ldapConfig.BindUser,
		bindPassword: ldapConfig.BindPassword,
		ldapServer:   ldapConfig.LDAPServer,
		ldapPort:     ldapConfig.LDAPPort,
		useTLS:       ldapConfig.UseTLS,
		startTLS:     ldapConfig.StartTLS,
	}
}

// Authenticate binds to the LDAP server with username and password
func (l *ldapService) Authenticate(username string, password string) (*models.User, error) {
	err := l.bind()
	if err != nil {
		return nil, err
	}
	defer l.conn.Close()

	searchRequest := ldap.NewSearchRequest(
		l.baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if len(sr.Entries) < 1 {
		err = fmt.Errorf("User %s does not exist", username)
	}

	userdn := sr.Entries[0].DN

	err = l.conn.Bind(userdn, password)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Username: username,
		Email:    "",
	}, nil
}

func (l *ldapService) bind() (err error) {
	if l.useTLS {
		l.conn, err = ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", l.ldapServer, l.ldapPort), &tls.Config{ServerName: l.ldapServer})
		if err != nil {
			return
		}
	} else {
		l.conn, err = ldap.Dial("tcp", fmt.Sprintf("%s:%d", l.ldapServer, l.ldapPort))
		if err != nil {
			return
		}

		if l.startTLS {
			err = l.conn.StartTLS(&tls.Config{ServerName: l.ldapServer, InsecureSkipVerify: true})
			if err != nil {
				return
			}
		}
	}

	err = l.conn.Bind(l.bindUser, l.bindPassword)
	if err != nil {
		return
	}

	return
}
