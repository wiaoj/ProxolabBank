package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string
	FirstName    string
	LastName     string
	Age          byte
	Email        string
	PhoneNumber  string
	PasswordHash []byte

	Claims []*Claim `gorm:"many2many:users_claims;"`
}
