package models

import "gorm.io/gorm"

type CreditType struct {
	gorm.Model
	Description string
	// TimeOptionID uint
	TimeOptions []*TimeOption `gorm:"many2many:creditType_timeOption;"`
	Interests   []Interest
}
