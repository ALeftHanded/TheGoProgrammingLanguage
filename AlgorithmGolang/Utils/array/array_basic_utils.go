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
