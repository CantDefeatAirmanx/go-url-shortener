package runes

import (
	"math"
	"math/rand/v2"
)

func PickRuneBetween(startRune rune, endRune rune) rune {
	randN := rand.Float64()

	randRune := rune(math.Round((randN * (float64(endRune) - float64(startRune))) + float64(startRune)))

	return randRune
}
