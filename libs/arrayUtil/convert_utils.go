package arrayUtil

// ConvertAnySliceToInterface  Helper function to convert []any to []interface{}
func ConvertAnySliceToInterface(slice []any) []interface{} {
	var iSlice = make([]interface{}, len(slice))
	for i, v := range slice {
		iSlice[i] = v
	}
	return iSlice
}

// ConvertIntSliceToInterface  Helper function to convert []int to []interface{}
func ConvertIntSliceToInterface(slice []int) []interface{} {
	var iSlice = make([]interface{}, len(slice))
	for i, v := range slice {
		iSlice[i] = v
	}
	return iSlice
}

// ConvertStringSliceToInterface  Helper function to convert []string to []interface{}
func ConvertStringSliceToInterface(slice []string) []interface{} {
	var iSlice = make([]interface{}, len(slice))
	for i, v := range slice {
		iSlice[i] = v
	}
	return iSlice
}
