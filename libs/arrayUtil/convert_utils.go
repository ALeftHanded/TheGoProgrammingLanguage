package arrayUtil

func ConvertIntToInterface(arr []int) []interface{} {
	var result []interface{}
	for _, value := range arr {
		result = append(result, value)
	}
	return result
}
