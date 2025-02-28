package auth

import (
	"fmt"

	auth "github.com/personal-project/zentio/internal/repository"
	"github.com/personal-project/zentio/internal/schema"
	"github.com/personal-project/zentio/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *auth.AuthRepository
}

func NewAuthService(repo *auth.AuthRepository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) SignUpUser(user *schema.User) error {
	// checking for the existing user by email
	err := s.Repo.CheckIfUserExistsWithEmail(user.Email)
	if err != nil {
		return err
	}

	// checking for the existing user by username
	err = s.Repo.CheckIfUserExistsWithUsername(user.Username)
	if err != nil {
		return err
	}

	// hashing password for new user
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	// creating a new user with hashed password
	if err := s.Repo.CreateUser(user); err != nil {
		return fmt.Errorf("failed to create user %w", err)
	}

	return nil
}

func (s *AuthService) SignInUser(email string, password string) (*schema.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	return user, nil
}