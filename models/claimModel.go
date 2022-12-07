package models

import "gorm.io/gorm"

type Claim struct {
	gorm.Model
	Name  string
	Level int
	Users []*User `gorm:"many2many:users_claims;"`
}

const (
	AdminClaimLevel = 1
	UserClaimLevel  = 2
)
