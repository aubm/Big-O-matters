package bigo

import (
	"math/rand"
	"time"
)

func newIntArray(length int) []int {
	array := make([]int, length)
	for i := 0; i < len(array); i++ {
		array[i] = random(10, 100)
	}
	return array
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
