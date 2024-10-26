package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	vendors []*Vendor `json:"vendors"`
}
