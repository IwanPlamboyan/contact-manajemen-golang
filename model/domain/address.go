package domain

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street     string `gorm:"column:street"`
	City       string `gorm:"column:city"`
	Province   string `gorm:"column:province"`
	Country    string `gorm:"column:country"`
	PostalCode string `gorm:"column:postal_code"`
	ContactID  uint   `gorm:"column:contact_id"`
}

func (Address) TableName() string {
	return "addresses"
}