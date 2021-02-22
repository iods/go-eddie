package parse

import "github.com/iods/go-eddie/internal/db/schema"


func Emojis(e []string) []schema.Emoji {
	var emojis []schema.Emoji

	l := len(e)
	for i := 0; i < l; i++ {
		emoji := []schema.Emoji{{Name: e[i]}}
		emojis = append(emojis, emoji...)
	}

	return emojis
}

func Tags(t []string) []schema.Tag {
	var tags []schema.Tag

	l := len(t)
	for i := 0; i < l; i++ {
		tag := []schema.Tag{{Name: t[i]}}
		tags = append(tags, tag...)
	}

	return tags
}