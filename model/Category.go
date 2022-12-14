package model

import (
	"golang.org/x/text/currency"
	"time"
)

type Category struct {
	ID        currency.Unit `json:"id" gorm:"primary_key"`
	Name      string        `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt time.Time     `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"type:timestamp"`
}
