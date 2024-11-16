package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`
    Email    string `json:"email"`
}