package gommon

import (
	"fmt"
	"github.com/hamidteimouri/gommon/htregex"
	"strconv"
	"strings"
)

// MakeMaskEmailAndDomain removes email's characters.
// for example converts 'username@gmail.com' to 'use***@gm**.com'
func MakeMaskEmailAndDomain(email string) string {
	if email == "" {
		return ""
	}
	if !htregex.IsEmail(email) {
		return MakeMaskUsername(email)
	}
	split := strings.Split(email, "@")

	hiddenEmail := ""
	for i := 0; i < len(split[0]); i++ {
		if i >= 3 {
			if i <= 4 {
				hiddenEmail += "*"
			}
		} else {
			hiddenEmail += string(split[0][i])
		}
	}

	splitDomain := strings.Split(split[1], ".")
	hiddenDomain := ""
	if len(splitDomain) == 2 {
		for i := 0; i < len(splitDomain[0]); i++ {
			if i >= 1 {
				if i <= 2 {
					hiddenDomain += "*"
				}
			} else {
				hiddenDomain += string(splitDomain[0][i])
			}
		}
		hiddenDomain += "." + splitDomain[1]
	} else {
		hiddenDomain = split[1]
	}

	hiddenEmail += "@" + hiddenDomain
	return hiddenEmail
}

// MakeMaskEmail removes email's characters.
// for example converts 'username@gmail.com' to 'us**@gmail.com'
func MakeMaskEmail(email string) string {
	if email == "" {
		return ""
	}
	if !htregex.IsEmail(email) {
		return MakeMaskUsername(email)
	}
	split := strings.Split(email, "@")

	hiddenEmail := ""
	for i := 0; i < len(split[0]); i++ {
		if i >= 2 {
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

// MakeMaskUsername removes username's characters.
// for example converts 'username' to 'us***me'
func MakeMaskUsername(username string) string {
	if username == "" {
		return ""
	}

	masked := ""
	for i := 0; i < len(username)-2; i++ {
		if i >= 2 {
			if i <= 4 {
				masked += "*"
			}
		} else {
			masked += string(username[i])
		}
	}
	masked += username[len(username)-2:]
	return masked
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

// Slugify converts a string to a slug.
// for example converts "this is test" to "this-is-test"
func Slugify(str string) string {
	return strings.Replace(str, " ", "-", -1)
}
