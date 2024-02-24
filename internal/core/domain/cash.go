package domain

import "gorm.io/gorm"

type Cash struct {
	gorm.Model
	Date        string   `json:"date"`
	Description string   `json:"description"`
	Value       float64  `json:"value"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category"`
	Source      string   `json:"source"`
}
