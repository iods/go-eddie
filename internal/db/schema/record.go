package schema

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model

	Length    int       `json:"length" fake:"{randomstring:[6:00,7:00,8:00,10:00]}"`
	Emojis	  []Emoji   `gorm:"many2many:record_emojis;" json:"emojis"`
	Frequency int       `json:"frequency"`
	From      time.Time `json:"from"`
	Important bool      `json:"important" fake:"{bool}"`
	Location  string    `json:"location"`
	Quality   int       `json:"quality" fake:"{number:1,10}"`
	Tags      []Tag     `gorm:"many2many:record_tags;" json:"tags"`
	To        time.Time `json:"to"`
	Time	  time.Time `json:"time"`
	Total     int       `json:"total"`
	Type      string    `json:"type"`
}

type Emoji struct {
	Name string `gorm:"primary_key" json:"emoji_name" fake:"{emojitag}"`
}

type Tag struct {
	Name string `gorm:"primary_key" json:"tag_name" fake:"{adjective}"`
}