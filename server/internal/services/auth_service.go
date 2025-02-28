package auth

import (
	"fmt"

	auth "github.com/personal-project/zentio/internal/repository"
	"github.com/personal-project/zentio/internal/schema"
	"github.com/personal-project/zentio/internal/utils"
)

type AuthService struct {
	Repo *auth.AuthRepository
}

func NewAuthService(repo *auth.AuthRepository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) SignUpUser(user *schema.User) (string, error) {
	// checking for the existing user
	existingUser, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if existingUser != nil {
		return "", fmt.Errorf("user with %s email already exists", existingUser.Email)
	}

	// hashing password for new user
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	// creating a new user with hashed password
	if err := s.Repo.CreateUser(user); err != nil {
		return "", fmt.Errorf("failed to create user %w", err)
	}

	// generating a jwt token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
