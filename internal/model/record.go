package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model

	Duration  uint    `json:"duration"`
	Emojis	  []Emoji `gorm:"many2many:record_emojis;" json:"record_emojis"`
	EndTime   string  `json:"end_time"`
	Important bool    `json:"important"`
	Location  string  `json:"location"`
	Quality   uint    `json:"quality"`
	StartTime string  `json:"start_time"`
	Tags      []Tag   `gorm:"many2many:record_tags;" json:"record_tags"`
	Total     string  `json:"total"`
	Type      string  `json:"type"`
}

type Emoji struct {
	Name string `gorm:"primary_key" json:"emoji_name"`
}

type Tag struct {
	Name string `gorm:"primary_key" json:"tag_name"`
}