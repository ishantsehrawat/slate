package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	oauth2 "golang.org/x/oauth2" // alias to 'oauth2'
	"golang.org/x/oauth2/google"
	googleOauth2 "google.golang.org/api/oauth2/v2" // alias to 'googleOauth2'

	"github.com/ishant/slate-backend/config"
	"github.com/ishant/slate-backend/internal/models"
)

var GoogleOauthConfig = &oauth2.Config{
    ClientID:     config.GoogleClientID,
    ClientSecret: config.GoogleClientSecret,
    RedirectURL:  config.GoogleRedirectURL,
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
    Endpoint:     google.Endpoint,
}

// GenerateJWT generates a signed JWT token for a user
func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(72 * time.Hour).Unix(), // token expires in 3 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

// ExchangeCodeForUser fetches user info from Google and returns User model data
func ExchangeCodeForUser(ctx context.Context, code string) (*models.User, error) {

	fmt.Println("re::", config.GoogleClientID, config.GoogleRedirectURL)
	token, err := GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := GoogleOauthConfig.Client(ctx, token)
	oauth2Service, err := googleOauth2.New(client)
	if err != nil {
		return nil, err
	}


	userinfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}

	if userinfo.Email == "" || userinfo.Id == "" {
		return nil, errors.New("invalid user info from Google")
	}

	return &models.User{
		GoogleID: userinfo.Id,
		Email:    userinfo.Email,
		Name:     userinfo.Name,
	}, nil
}
