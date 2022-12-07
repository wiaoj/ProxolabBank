package models

import "gorm.io/gorm"

//faiz
type Interest struct {
	gorm.Model
	BankID       uint
	Bank         Bank
	Interest     float32
	TimeOptionID uint
	TimeOption   TimeOption
	CreditTypeID uint
	CreditType   CreditType
}
