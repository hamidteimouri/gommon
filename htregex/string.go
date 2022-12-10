package htregex

import (
	"fmt"
	"strconv"
	"strings"
)

// MakeEmailInvisible remove email's characters.
// for example converts username@gmail.com to use***@gmail.com
func MakeEmailInvisible(email string) string {
	if email == "" {
		return ""
	}
	if !IsEmail(email) {
		return MakeUsernameInvisible(email)
	}
	split := strings.Split(email, "@")

	hiddenEmail := ""
	for i := 0; i < len(split[0]); i++ {
		if i >= 3 {
			if i <= 5 {
				hiddenEmail += "*"
			}
		} else {
			hiddenEmail += string(split[0][i])
		}
	}

	hiddenEmail += "@" + split[1]
	return hiddenEmail
}

// MakeUsernameInvisible remove username's characters.
// for example converts username to us***me
func MakeUsernameInvisible(username string) string {
	if username == "" {
		return ""
	}

	hiddenEmail := ""
	for i := 0; i < len(username)-2; i++ {
		if i >= 2 {
			if i <= 4 {
				hiddenEmail += "*"
			}
		} else {
			hiddenEmail += string(username[i])
		}
	}
	hiddenEmail += username[len(username)-2:]
	return hiddenEmail
}

// RemoveThousandsSeparator removes (,) from the number.
// for example converts 1,009 to 1009
func RemoveThousandsSeparator(number string) (float64, error) {
	number = strings.Replace(number, ",", "", -1)
	unFormatted, err := strconv.ParseFloat(number, 32)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return 0, err
	}
	return unFormatted, nil
}
