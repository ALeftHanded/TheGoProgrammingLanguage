package Array

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
