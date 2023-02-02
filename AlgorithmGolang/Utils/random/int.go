package random

import "math/rand"

func GenRandomNums(minLen, maxLen, minElement, maxElement int) []int {
	var res []int
	ranLen := rand.Intn(maxLen-minLen) + minLen
	for i := 0; i < ranLen; i++ {
		randNum := rand.Intn(maxElement-minElement) + minElement
		res = append(res, randNum)
	}
	return res
}
