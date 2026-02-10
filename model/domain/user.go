package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;unique;not null"`
	Password string `gorm:"column:password;not null"`
	Name     string `gorm:"column:name;not null"`
	Contacts []Contact `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}