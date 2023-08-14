package arrayUtil

func InIntArray(key int, array []int) bool {
	if array == nil {
		return false
	}
	for _, value := range array {
		if value == key {
			return true
		}
	}
	return false
}

func Deduplicate(array []int) []int {
	if len(array) == 0 {
		return array
	}
	seen, deduplicated := make(map[int]bool), make([]int, 0, len(array))
	for _, val := range array {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			deduplicated = append(deduplicated, val)
		}
	}
	return deduplicated
}

// MaxInt is a function which returns the maximum of all the integers provided as arguments.
func MaxInt(values ...int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// MinInt is a function which returns the minimum of all the integers provided as arguments.
func MinInt(values ...int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
