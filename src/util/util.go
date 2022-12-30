package util

import (
	"math/rand"
	"time"
)

func GetRandValue(items []string) string {
	rand.Seed(time.Now().UnixNano())
	return items[rand.Intn(len(items))]
}
