package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(40);unique" json:"username,omitempty" binding:"required"`
	Password string `gorm:"size:255" json:"password,omitempty" binding:"required"`
}
