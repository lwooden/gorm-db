package Models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
