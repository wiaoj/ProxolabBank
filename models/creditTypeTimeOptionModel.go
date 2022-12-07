package models

type CreditTypeTimeOption struct {
	CreditTypeID int `gorm:"primaryKey"`
	TimeOptionID int `gorm:"primaryKey"`
}
