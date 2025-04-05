package gommon

import (
	"fmt"
	"strconv"
)

const (
	TimezoneTehran = "Asia/Tehran"
)

// PersianDateSplitter splits a number into x, y, z parts based on specific rules.
// for example : '13690516' => 1369 , 05 , 16
func PersianDateSplitter(input string) (x, y, z int, err error) {
	// Convert the input number to a string
	//str := fmt.Sprintf("%08d", input) // Ensures it's padded to 8 digits if needed
	str := input // Ensures it's padded to 8 digits if needed

	// Ensure the string length is exactly 8
	if len(str) != 8 {
		return 0, 0, 0, fmt.Errorf("input must be an 8-digit number")
	}

	// Parse the parts
	x, err = strconv.Atoi(str[:4]) // First 4 digits
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse x: %v", err)
	}
	y, err = strconv.Atoi(str[4:6]) // Next 2 digits
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse y: %v", err)
	}
	z, err = strconv.Atoi(str[6:]) // Last 2 digits
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse z: %v", err)
	}

	return x, y, z, nil
}
