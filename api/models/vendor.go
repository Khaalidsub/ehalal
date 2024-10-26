package models

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	CategoryID  string `json:"category_id"`
}
