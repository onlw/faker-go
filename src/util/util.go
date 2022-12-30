package util

import (
	"math/rand"
	"time"
)

const hashtag = '#'

func GetRandValue(items []string) string {
	rand.Seed(time.Now().UnixNano())
	return items[rand.Intn(len(items))]
}

func RandIntRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	// If they pass in the same number, just return that number
	if min == max {
		return min
	}

	// If they pass in a min that is bigger than max, swap them
	if min > max {
		ogmin := min
		min = max
		max = ogmin
	}

	// Figure out if the min/max numbers calculation
	// would cause a panic in the Int63() function.
	if max-min+1 > 0 {
		return min + int(rand.Int63n(int64(max-min+1)))
	}

	// Loop through the range until we find a number that fits
	for {
		v := int(rand.Uint64())
		if (v >= min) && (v <= max) {
			return v
		}
	}
}

// ReplaceWithNumbers  Replace # with numbers
func ReplaceWithNumbers(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = byte(randDigit())
		}
	}
	if bytestr[0] == '0' {
		bytestr[0] = byte(rand.Intn(8)+1) + '0'
	}

	return string(bytestr)
}

// Generate random ASCII digit
func randDigit() rune {
	return rune(byte(rand.Intn(10)) + '0')
}
