package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Lastname string `json:"lastname" gorm:"not null"`
	Username string `gorm:"unique;size:6;not null" json:"username"`
	Email    string `gorm:"unique;size:10;not null" json:"email"`
	Password string `gorm:"unique;size:6;not null" json:"password"`
	Gender   string `json:"gender"`
	Role     string `json:"role"`
}
