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

// InArray is a function like Contains
func InArray(arr []string, str string) bool {
	for _, v := range arr {
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

func ConvertMapToSlice[T any](m map[any]T) []T {
	s := make([]T, 0)
	for _, v := range m {
		s = append(s, v)
	}

	return s
}

func CloneMap[K comparable, V any](m map[K]V) map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		n[k] = v
	}

	return n
}

func CloneSlice[V any](s []V) []V {
	t := make([]V, 0)
	for _, v := range s {
		t = append(t, v)
	}

	return t
}
