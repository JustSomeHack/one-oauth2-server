package authenticationservice

import (
	"fmt"
	"os"
	"time"

	"github.com/JustSomeHack/one-oauth2-server/models/mongoconfig"

	"github.com/JustSomeHack/one-oauth2-server/models/ldapconfig"
	"github.com/JustSomeHack/one-oauth2-server/services/ldapservice"
	"github.com/JustSomeHack/one-oauth2-server/services/mongoservice"

	"github.com/JustSomeHack/one-oauth2-server/models"
	"github.com/dgrijalva/jwt-go"
)

// AuthenticationService to authenticate user
type AuthenticationService interface {
	Authorize(username string, password string) (map[string]string, error)
	Validate(bearerToken string) (interface{}, error)
}

type authenticationService struct {
	authType   string
	userSecret []byte
}

// NewAuthenticationService gets a reference to AuthenticationService
func NewAuthenticationService() AuthenticationService {
	return &authenticationService{
		authType:   os.Getenv("AUTH_TYPE"),
		userSecret: []byte(os.Getenv("USER_SECRET")),
	}
}

// Authorize authenticates a user against a service
func (a *authenticationService) Authorize(username string, password string) (userMap map[string]string, err error) {
	user := new(models.User)

	switch a.authType {
	case "ldap":
		ldapConfig := ldapconfig.LoadLDAPConfig()
		ldap := ldapservice.NewLDAPService(ldapConfig)
		user, err = ldap.Authenticate(username, password)
		break
	case "mongo":
		mongoConfig := mongoconfig.LoadMongoConfig()
		mongo := mongoservice.NewMongoService(mongoConfig)
		user, err = mongo.Authenticate(username, password)
		break
	}

	if err != nil {
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "one-oauth2-server",
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
		"username": user.Username,
		"email":    user.Email,
	})

	tokenString, err := token.SignedString(a.userSecret)
	if err != nil {
		return
	}

	userMap = map[string]string{
		"username": user.Username,
		"email":    user.Email,
		"token":    tokenString,
	}
	return
}

func (a *authenticationService) Validate(bearerToken string) (interface{}, error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return a.userSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return token.Claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
