package models

import "gorm.io/gorm"

type SearchHistory struct {
	gorm.Model
	User   User   `gorm:"ForeignKey:ID" binding:"required"`
	Search string `gorm:"size:255" json:"password,omitempty" binding:"required"`
}
