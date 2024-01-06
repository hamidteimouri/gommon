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
// for example converts "this is a test" to "this-is-a-test"
func Slugify(str string) string {
	return strings.Replace(str, " ", "-", -1)
}

// ConvertArabicAlphabetToPersian replace arabic character with persian equivalent
func ConvertArabicAlphabetToPersian(str string) string {
	table := map[string]string{"ك": "ک", "ي": "ی"}
	for ar, fa := range table {
		str = strings.ReplaceAll(str, ar, fa)
	}

	return str
}

// ConvertNumberArabicToEnglish replace all arabic numbers to english equivalent
func ConvertNumberArabicToEnglish(num string) string {
	table := map[string]string{"٠": "0", "١": "1", "٢": "2",
		"٣": "3", "٤": "4",
		"٥": "5", "٦": "6", "٧": "7",
		"٨": "8", "٩": "9",
	}
	for per, eng := range table {
		num = strings.ReplaceAll(num, per, eng)
	}

	return num
}

// ConvertNumberPersianToEnglish replace persian number with english equivalent
func ConvertNumberPersianToEnglish(num string) string {
	table := map[string]string{"۰": "0", "۱": "1", "۲": "2", "۳": "3",
		"۴": "4", "۵": "5", "۶": "6", "۷": "7",
		"۸": "8", "۹": "9",
	}
	for per, eng := range table {
		num = strings.ReplaceAll(num, per, eng)
	}

	return num
}
