package domain

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName string `json:"full_name"`
	Username string `gorm:"type:VARCHAR(50); uniqueIndex; NOT NULL" json:"username" `
	Email    string `gorm:"type:VARCHAR(50); uniqueIndex; NOT NULL" json:"email" `
	Password string `json:"password"`
	Gender   bool   `json:"gender"`
	Picture  string `json:"picture"`
	Role     string `gorm:"NOT NULL" json:"role"`
}

type UserRegister struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   bool   `json:"gender"`
	Role     string `binding:"required" json:"role"`
}

type UserLogin struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Picture  string `json:"picture"`
	Gender   string `json:"gender"`
}
