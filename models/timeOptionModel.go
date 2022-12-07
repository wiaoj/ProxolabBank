package models

import "gorm.io/gorm"

type TimeOption struct {
	gorm.Model
	Description string
	// CreditTypeID uint
	CreditTypes []*CreditType `gorm:"many2many:creditType_timeOption;"`
	Interests   []Interest
}
