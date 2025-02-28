package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=20"`
	Name     string `gorm:"not null" json:"name" validate:"required,min=3,max=20"`
	Email    string `gorm:"unique;not null" json:"email,omitempty" validate:"required,email"`
	Password string `gorm:"not null" json:"password,omitempty" validate:"required,min=8,max=20"`
}

type SignInUser struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=20"`
}

type SignedUser struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Name     string `json:"name" validate:"required,min=3,max=20"`
}
