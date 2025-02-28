package auth

import (
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

func (r *AuthRepository) GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	err := r.DB.Where("email = ?", email).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
