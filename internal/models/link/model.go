package link

import (
	"math/rand"
	"url_shortener/pkg/runes"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {

	return &Link{Url: url, Hash: randStringRunes(10)}
}

func randStringRunes(n int) string {
	runeSlice := make([]rune, n)

	for index := range runeSlice {

		startRune := 'a'
		endRune := 'z'

		randN := rand.Float64()
		if randN > 0.5 {
			startRune = 'A'
			endRune = 'Z'
		}

		runeSlice[index] = runes.PickRuneBetween(startRune, endRune)
	}

	return string(runeSlice)
}
