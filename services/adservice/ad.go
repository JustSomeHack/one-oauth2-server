package adservice

import (
	"github.com/JustSomeHack/one-oauth2-server/models"
)

// ADService service to connect to Active Directory server
type ADService interface {
	Authenticate(username string, password string) (*models.User, error)
}

type adService struct {
	baseDN       string
	bindUser     string
	bindPassword string
	ldapServer   string
	ldapPort     int
	useTLS       bool
}

// NewADService gets a new reference to ADService
func NewADService() ADService {
	return &adService{}
}

// Authenticate binds to the Active Directory server with username and password
func (a *adService) Authenticate(username string, password string) (*models.User, error) {

	return &models.User{
		Username: username,
		Email:    "",
	}, nil
}
