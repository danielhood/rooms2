package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/danielhood/rooms2/api/models"
	"github.com/danielhood/rooms2/api/repo"
)

// Set our secret.
var mySigningKey = []byte("jlrew03h3@4")

// Token defines a token for our application
type Token string

// TokenService provides a token
type TokenService interface {
	ProcessUserLogin(username string, password string) (string, error)
}

type tokenService struct {
	userRepo repo.UserRepo
}

type userClaims struct {
	IsAdmin  bool   `json:"isadmin"`
	AuthType string `json:"authtype"`
	jwt.StandardClaims
}

// NewTokenService creates a new UserService
func NewTokenService(
	userRepo repo.UserRepo,
) TokenService {
	return &tokenService{
		userRepo: userRepo,
	}
}

func (s *tokenService) ProcessUserLogin(username string, password string) (string, error) {
	log.Print("Request User: ", username)

	user, err := s.userRepo.GetByUsername(username)
	//user, err := s.getByUsername(username)

	if err != nil {
		log.Print("Error retrieving username: ", err)
		return "", err
	}

	if password != user.Password {
		log.Print("Invalid password")
		return "", errors.New("invalid password")
	}

	if !user.IsEnabled {
		return "", errors.New("user not enabled")
	}

	return s.getUserToken(user)
}

// GetUserToken retrieves a token for a user
func (s *tokenService) getUserToken(u *models.User) (string, error) {
	// Set token claims
	claims := userClaims{
		u.HasRole(models.AdministratorRole),
		"user",
		jwt.StandardClaims{
			Subject:   u.Username,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "token-service",
			Id:        uuid.NewString(),
		},
	}

	tokenString, _ := s.createToken(claims)

	fmt.Printf("Generated token for %v: %v\n", u.Username, tokenString)

	return tokenString, nil
}

func (s *tokenService) createToken(claims userClaims) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}
