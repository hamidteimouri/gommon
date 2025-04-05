package gommon

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// MakeMaskEmailAndDomain removes email's characters.
// for example converts 'username@gmail.com' to 'use***@gm**.com'
func MakeMaskEmailAndDomain(email string) string {
	if email == "" {
		return ""
	}
	if !IsEmail(email) {
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
	if !IsEmail(email) {
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

func TrimSpacesTabsNewlines(inputString string) string {
	// Replace all spaces, tabs, and newlines with an empty string
	trimmedString := strings.ReplaceAll(inputString, " ", "")
	trimmedString = strings.ReplaceAll(trimmedString, "\t", "")
	trimmedString = strings.ReplaceAll(trimmedString, "\n", "")
	trimmedString = strings.ReplaceAll(trimmedString, "\r", "")

	return trimmedString
}

func CutPrecision(input string, precision int) string {
	// Split the string into the integer and fractional parts
	parts := strings.Split(input, ".")
	if len(parts) == 2 {
		// Cut the fractional part to the required precision
		if len(parts[1]) > precision {
			parts[1] = parts[1][:precision]
		}
		// Recombine the integer and cut fractional parts
		return parts[0] + "." + parts[1]
	}

	input = strings.TrimRight(input, ".")

	// If no fractional part, return the input as-is
	return input
}

// FormatNumber adds commas as thousands separators to a number string.
func FormatNumber(input string) string {
	// Split the input into the integer and fractional parts
	parts := strings.Split(input, ".")
	integerPart := parts[0]

	// Handle the integer part
	var buffer bytes.Buffer
	length := len(integerPart)
	for i, char := range integerPart {
		buffer.WriteRune(char)
		if (length-i-1)%3 == 0 && i != length-1 {
			buffer.WriteRune(',')
		}
	}

	// Append the fractional part if it exists
	if len(parts) > 1 {
		buffer.WriteString(".")
		buffer.WriteString(parts[1])
	}

	return buffer.String()
}
