package schema

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model

	Length    int       `json:"length" fake:"{number:3,8}"`
	Emojis	  []Emoji   `gorm:"many2many:record_emojis;" json:"emojis"`
	Frequency int       `json:"frequency" fake:"{number:1,8}"`
	From      time.Time `json:"from"`
	Important bool      `json:"important" fake:"{bool}"`
	Location  string    `json:"location" fake:"{randomstring:[office,couch,bedroom]}"`
	Quality   int       `json:"quality" fake:"{number:1,10}"`
	Stress    int       `json:"stress" fake:"{number:1,10}"`
	Tags      []Tag     `gorm:"many2many:record_tags;" json:"tags"`
	To        time.Time `json:"to"`
	Time	  time.Time `json:"time"`
	Total     float64   `json:"total" fake:"{number:180,200}"`
	Type      string    `json:"type"`
}

type Emoji struct {
	Name string `gorm:"primary_key" json:"emoji_name" fake:"{randomstring:[joyful,happy,grateful,sad,anxious,bored,loving,loved,stressed,focused,frustrated,hopeful]}"`
}

type Tag struct {
	Name string `gorm:"primary_key" json:"tag_name" fake:"{adjective}"`
}