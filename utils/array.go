package utils

func InArray(needle interface{}, haystack []interface{}) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}
	return false
}
