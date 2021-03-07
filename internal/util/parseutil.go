package parse

import (
	"github.com/iods/go-eddie/internal/db/schema"
)

// Emojis Returns a slice of emojis for construction of a record.
func Emojis(e []string) []schema.Emoji {
	var emojis []schema.Emoji
	l := len(e)
	for i := 0; i < l; i++ {
		emoji := []schema.Emoji{{Name: e[i]}}
		emojis = append(emojis, emoji...)
	}
	return emojis
}

// Returns a slice of Tags for construction of a record.
func Tags(t []string) []schema.Tag {
	var tags []schema.Tag
	l := len(t)
	for i := 0; i < l; i++ {
		tag := []schema.Tag{{Name: t[i]}}
		tags = append(tags, tag...)
	}
	return tags
}