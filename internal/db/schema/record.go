package schema

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model

	Length    int   `json:"length"`
	Emojis	  []Emoji `gorm:"many2many:record_emojis;" json:"emojis"`
	Frequency int     `json:"frequency"`
	From      time.Time  `json:"from"`
	Important bool    `json:"important"`
	Location  string  `json:"location"`
	Quality   int   `json:"quality"`
	Tags      []Tag   `gorm:"many2many:record_tags;" json:"tags"`
	To        time.Time  `json:"to"`
	Time	  time.Time `json:"time"`
	Total     int   `json:"total"`
	Type      string  `json:"type"`
}

type Emoji struct {
	Name string `gorm:"primary_key" json:"emoji_name"`
}

type Tag struct {
	Name string `gorm:"primary_key" json:"tag_name"`
}