package gommon

import (
	"github.com/hamidteimouri/gommon/htregex"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const maxURLRuneCount = 2083
const minURLRuneCount = 3

const (
	IP string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`

	URLSchema    string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername  string = `(\S+(:\S*)?@)`
	URLPath      string = `((\/|\?|#)[^\s]*)`
	URLPort      string = `(:(\d{1,5}))`
	URLIP        string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	URLSubdomain string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	URL                 = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
)

var (
	rgxUrl = regexp.MustCompile(URL)
)

// IsUTFLetter checks if the string contains only unicode letter characters.
// Similar to IsAlpha but for all languages. Empty string is valid.
func IsUTFLetter(str string) bool {
	if IsNull(str) {
		return true
	}

	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true

}

// IsUTFLetterNumeric checks if the string contains only unicode letters and numbers. Empty string is valid.
func IsUTFLetterNumeric(str string) bool {
	if IsNull(str) {
		return true
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) { //letters && numbers are ok
			return false
		}
	}
	return true

}

// IsUTFNumeric checks if the string contains only unicode numbers of any kind.
// Numbers can be 0-9 but also Fractions ¾,Roman Ⅸ and Hangzhou 〩. Empty string is valid.
func IsUTFNumeric(str string) bool {
	if IsNull(str) {
		return true
	}
	if strings.IndexAny(str, "+-") > 0 {
		return false
	}
	if len(str) > 1 {
		str = strings.TrimPrefix(str, "-")
		str = strings.TrimPrefix(str, "+")
	}
	for _, c := range str {
		if !unicode.IsNumber(c) { //numbers && minus sign are ok
			return false
		}
	}
	return true

}

// IsUTFDigit checks if the string contains only unicode radix-10 decimal digits. Empty string is valid.
func IsUTFDigit(str string) bool {
	if IsNull(str) {
		return true
	}
	if strings.IndexAny(str, "+-") > 0 {
		return false
	}
	if len(str) > 1 {
		str = strings.TrimPrefix(str, "-")
		str = strings.TrimPrefix(str, "+")
	}
	for _, c := range str {
		if !unicode.IsDigit(c) { //digits && minus sign are ok
			return false
		}
	}
	return true

}

// IsLowerCase checks if the string is lowercase. Empty string is valid.
func IsLowerCase(str string) bool {
	if IsNull(str) {
		return true
	}
	return str == strings.ToLower(str)
}

// IsUpperCase checks if the string is uppercase. Empty string is valid.
func IsUpperCase(str string) bool {
	if IsNull(str) {
		return true
	}
	return str == strings.ToUpper(str)
}

// IsNull checks if the string is null.
func IsNull(str string) bool {
	return len(str) == 0
}

// IsNotNull checks if the string is not null.
func IsNotNull(str string) bool {
	return !IsNull(str)
}

// IsByteLength checks if the string's length (in bytes) falls in a range.
func IsByteLength(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}

// IsURL checks if the string is a URL.
func IsURL(str string) bool {
	if str == "" || utf8.RuneCountInString(str) >= maxURLRuneCount || len(str) <= minURLRuneCount || strings.HasPrefix(str, ".") {
		return false
	}
	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used, so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return rgxUrl.MatchString(str)
}

// IsRequestURL checks if the string rawurl, assuming
// it was received in an HTTP request, is a valid
// URL confirm to RFC 3986
func IsRequestURL(rawUrl string) bool {
	url1, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return false //Couldn't even parse the rawurl
	}
	if len(url1.Scheme) == 0 {
		return false //No Scheme found
	}
	return true
}

// IsRequestURI checks if the string rawUrl, assuming
// it was received in an HTTP request, is an
// absolute URI or an absolute path.
func IsRequestURI(rawUrl string) bool {
	_, err := url.ParseRequestURI(rawUrl)
	return err == nil
}

func IsValidHttpsUrl(url string) bool {
	if !IsURL(url) {
		return false
	}
	if strings.Contains(url, "https://") {
		return true
	}
	return false
}

func IsEmailInDomain(email, domain string) bool {
	if !htregex.IsEmail(email) {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		// Invalid email format
		return false
	}

	// Compare the extracted domain with the desired domain
	return parts[1] == domain
}

// ValidateCreditCardNumber validate credit card number (persian one)
func ValidateCreditCardNumber(code string) bool {
	L := len(code)
	if L < 16 || parseInt(code[1:11]) == 0 || parseInt(code[10:16]) == 0 {
		return false
	}
	s := 0
	var k, d int
	for i := 0; i < 16; i++ {
		k = 1
		if i%2 == 0 {
			k = 2
		}
		d, _ = strconv.Atoi(string(code[i]))
		d *= k
		if d > 9 {
			d -= 9
		}
		s += d
	}
	return s%10 == 0
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
