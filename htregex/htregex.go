package htregex

/*
*  Used https://github.com/mingrammer/commonregex/blob/master/commonregex.go
 */

import (
	"regexp"
	"strings"
)

const (
	EmailsPattern = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
	EmailPattern  = `^(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)$`
	IPv4Pattern   = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
	IPv6Pattern   = `(?:(?:(?:[0-9A-Fa-f]{1,4}:){7}(?:[0-9A-Fa-f]{1,4}|:))|(?:(?:[0-9A-Fa-f]{1,4}:){6}(?::[0-9A-Fa-f]{1,4}|(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){5}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,2})|:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){4}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,3})|(?:(?::[0-9A-Fa-f]{1,4})?:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){3}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,4})|(?:(?::[0-9A-Fa-f]{1,4}){0,2}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){2}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,5})|(?:(?::[0-9A-Fa-f]{1,4}){0,3}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){1}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,6})|(?:(?::[0-9A-Fa-f]{1,4}){0,4}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?::(?:(?:(?::[0-9A-Fa-f]{1,4}){1,7})|(?:(?::[0-9A-Fa-f]{1,4}){0,5}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(?:%.+)?\s*`
	IPPattern     = IPv4Pattern + `|` + IPv6Pattern
	IBANPattern   = `[A-Z]{2}\d{2}[A-Z0-9]{4}\d{7}([A-Z\d]?){0,16}`
)

var (
	EmailRegex  = regexp.MustCompile(EmailPattern)
	EmailsRegex = regexp.MustCompile(EmailsPattern)
	IPv4Regex   = regexp.MustCompile(IPv4Pattern)
	IPv6Regex   = regexp.MustCompile(IPv6Pattern)
	IPRegex     = regexp.MustCompile(IPPattern)
	IBANRegex   = regexp.MustCompile(IBANPattern)
)

func match(text string, regex *regexp.Regexp) []string {
	p := regex.FindAllString(text, -1)
	return p
}

// Emails finds all email strings
func Emails(text string) []string {
	return match(text, EmailsRegex)
}

// IsEmail check regex of an email
func IsEmail(text string) bool {
	return EmailRegex.MatchString(text)
}

// IsEmailUsingTrimLR check regex of an email that has been trim
func IsEmailUsingTrimLR(text string) bool {
	return EmailRegex.MatchString(LRTrimSpace(text))
}

// IPv4s finds all IPv4 addresses
func IPv4s(text string) []string {
	return match(text, IPv4Regex)
}

// IPv6s finds all IPv6 addresses
func IPv6s(text string) []string {
	return match(text, IPv6Regex)
}

// IPs finds all IP addresses (both IPv4 and IPv6)
func IPs(text string) []string {
	return match(text, IPRegex)
}

// IBANs finds all IBAN strings
func IBANs(text string) []string {
	return match(text, IBANRegex)
}

// LRTrimSpace trims left and write spaces
func LRTrimSpace(text string) string {
	text = strings.TrimLeft(text, " ")
	return strings.TrimRight(text, " ")
}
