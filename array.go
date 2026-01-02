package gommon

import (
	"reflect"
	"strings"
)

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

func ConvertQueryStringToArray(query string) []string {
	if query == "" {
		return nil
	}
	query = strings.TrimLeft(query, "[")
	query = strings.TrimRight(query, "]")
	query = strings.TrimRight(query, ",")
	query = strings.TrimLeft(query, ",")
	splitQuery := strings.Split(query, ",")
	return splitQuery
}

func ConvertStringToArray(query string) []string {
	if query == "" {
		return nil
	}
	query = strings.TrimLeft(query, "[")
	query = strings.TrimRight(query, "]")
	query = strings.TrimRight(query, ",")
	query = strings.TrimLeft(query, ",")
	splitQuery := strings.Split(query, ",")
	return splitQuery
}

func IsSlice(v interface{}) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsArray(v interface{}) bool {
	return IsSlice(v)
}

func RemoveDuplicates[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	var list []T
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
