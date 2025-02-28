package auth

import (
	"fmt"

	"github.com/personal-project/zentio/internal/schema"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) CreateUser(user *schema.User) error {
	return r.DB.Create(user).Error
}

func (r *AuthRepository) CheckIfUserExistsWithEmail(email string) error {
	var user schema.User
	err := r.DB.Where("email = ?", email).Take(&user).Error

	if err == nil {
		return fmt.Errorf("an account with the email '%s' already exists", email)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

func (r *AuthRepository) CheckIfUserExistsWithUsername(username string) error {
	var user schema.User

	err := r.DB.Where("username = ?", username).Take(&user).Error
	if err == nil {
		return fmt.Errorf("an account with the username '%s' already exists", username)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	err := r.DB.Where("email = ?", email).Take(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
