package gommon

// Contains check the string is in an array or not
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ArrayOverlap(mainArray, secondArray []string) bool {
	exists := make(map[string]bool)
	for _, value := range mainArray {
		exists[value] = true
	}
	for _, value := range secondArray {
		if !exists[value] {
			return false
		}
	}
	return true
}
