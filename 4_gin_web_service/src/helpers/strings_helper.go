package helpers

func ContainsString(srcStringArray []string, key string) bool {
	for _, value := range srcStringArray {
		if value == key {
			return true
		}
	}
	return false
}
