package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=20"`
	Name     string `gorm:"not null" json:"name" validate:"required,min=3,max=20"`
	Email    string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required,min=8,max=20"`
}
