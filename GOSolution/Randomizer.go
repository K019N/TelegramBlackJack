package main

import "math/rand"

type Randomizer struct{}

func randIntIn(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randPictureCard() int {
	return rand.Intn(3)
}
