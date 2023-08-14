package random

import (
	"math/rand"
	"time"
)

func GenRandomNums(minLen, maxLen, minElement, maxElement int) []int {
	// seed reset
	rand.Seed(time.Now().UnixNano())

	var res []int

	var ranLen int
	if maxLen > minLen {
		ranLen = rand.Intn(maxLen-minLen) + minLen
	} else {
		ranLen = minLen
	}
	for i := 0; i < ranLen; i++ {
		randNum := rand.Intn(maxElement-minElement) + minElement
		res = append(res, randNum)
	}
	return res
}

func GenRandomIncrIntList(minLen, maxLen, minElement, maxElement int) []int {
	// seed reset
	rand.Seed(time.Now().UnixNano())

	// Generate a random length for the list
	listLen := minLen + rand.Intn(maxLen-minLen+1) - 1

	// Check if the length of the list is too long and adjust it if necessary
	if maxElement-minElement < listLen {
		listLen = maxElement - minElement
	}

	// calculate random max step between two elements
	maxStep := (maxElement - minElement) / listLen
	// Initialize the list with the first element
	list := make([]int, 0, listLen)
	list = append(list, minElement)
	// Loop over the remaining elements and add them to the list
	if maxStep == 1 {
		for i := 1; i < listLen; i++ {
			list = append(list, list[i-1]+1)
		}
	} else {
		for i := 1; i < listLen; i++ {
			list = append(list, list[i-1]+rand.Intn(maxStep-1)+1)
		}
	}
	//fmt.Println(list)

	return list
}
