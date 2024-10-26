package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Rating  int16  `json:"rating"`
}
