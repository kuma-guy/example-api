package accounts

import (
	"net/http"

	"github.com/RichardKnop/recall/config"
	"github.com/RichardKnop/recall/oauth"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	// Exported methods
	GetConfig() *config.Config
	GetOauthService() oauth.ServiceInterface
	FindAccountByOauthClientID(oauthClientID uint) (*Account, error)
	FindAccountByID(accountID uint) (*Account, error)
	CreateAccount(name, description, key, secret, redirectURI string) (*Account, error)
	FindUserByOauthUserID(oauthUserID uint) (*User, error)
	FindUserByID(userID uint) (*User, error)
	FindUserByFacebookID(facebookID string) (*User, error)
	CreateUser(account *Account, userRequest *UserRequest) (*User, error)
	UpdateUser(user *User, userRequest *UserRequest) error
	CreateFacebookUser(account *Account, facebookID string, userRequest *UserRequest) (*User, error)
	CreateSuperuser(account *Account, email, password string) (*User, error)

	// Needed for the newRoutes to be able to register handlers
	createUserHandler(w http.ResponseWriter, r *http.Request)
	getMyUserHandler(w http.ResponseWriter, r *http.Request)
	updateUserHandler(w http.ResponseWriter, r *http.Request)
}
