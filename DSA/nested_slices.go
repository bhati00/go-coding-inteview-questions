package languagebasics

// flatten the nested slices into a single array

func flattenSlice(nums interface{}) []int {
	result := []int{}
	val, ok := nums.([]interface{})
	if ok {
		for _, elem := range val {
			switch v := elem.(type) {
			case int:
				result = append(result, v)
			case []interface{}:
				result = append(result, flattenSlice(v)...)
			}
		}
	}
	return result
}
