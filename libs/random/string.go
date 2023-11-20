package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomStringSequence(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}

func GenerateRandomStringSlice(minSliceLength, maxSliceLength, minStrLength, maxStrLength int) []string {
	// slice length in [minSliceLength, maxSliceLength]
	sliceLength := rand.Intn(maxSliceLength-minSliceLength+1) + minSliceLength

	res := make([]string, 0, sliceLength)

	for i := 0; i < sliceLength; i++ {
		// generate random string, length in [minStrLength, maxStrLength]
		randStrLen := rand.Intn(maxStrLength-minStrLength+1) + minStrLength
		tmpRandStr := GenerateRandomStringSequence(randStrLen)
		res = append(res, tmpRandStr)
	}
	return res
}
