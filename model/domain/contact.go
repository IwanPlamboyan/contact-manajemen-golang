package domain

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Phone     string `gorm:"column:phone"`
	UserID    uint   `gorm:"column:user_id"`
	User      User     `gorm:"foreignKey:user_id;references:id"`
	Addresses []Address `gorm:"foreignKey:contact_id;references:id"`
}

func (Contact) TableName() string {
	return "contacts"
}