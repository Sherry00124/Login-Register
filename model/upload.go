package model

import "gorm.io/gorm"

type Upload struct {
	gorm.Model
	Name    string `gorm:"type:varchar(20);not null"`
	Address string `gorm:"type:varchar(20);not null"`
}
