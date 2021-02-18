package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model

	Name string `json:"name"`
	Type string `json:"type"`
}