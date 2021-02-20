package schema

import "gorm.io/gorm"

type Record struct {
	gorm.Model

	Length    uint    `json:"length"`
	Emojis	  []Emoji `gorm:"many2many:record_emojis;" json:"emojis"`
	Frequency uint    `json:"frequency"`
	From      string  `json:"from"`
	Important bool    `json:"important"`
	Location  string  `json:"location"`
	Quality   uint    `json:"quality"`
	Tags      []Tag   `gorm:"many2many:record_tags;" json:"tags"`
	To        string  `json:"to"`
	Total     string  `json:"total"`
	Type      string  `json:"type"`
}

type Emoji struct {
	Name string `gorm:"primary_key" json:"emoji_name"`
}

type Tag struct {
	Name string `gorm:"primary_key" json:"tag_name"`
}