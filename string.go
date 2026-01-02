package gommon

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"
)

// SimilarityPercentage calculates the similarity percentage between two strings
func SimilarityPercentage(a, b string) float64 {
	if a != "" && b == "" {
		return 0
	}
	if b != "" && a == "" {
		return 0
	}
	a = strings.ReplaceAll(a, "\u200C", "") //  remove ZWNJ ( zero-width non-joiner )
	b = strings.ReplaceAll(b, "\u200C", "") //  remove ZWNJ ( zero-width non-joiner )
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	distance := levenshteinDistance(a, b)
	maxLen := maxItem(len(a), len(b))
	if maxLen == 0 {
		return 100.0
	}
	similarity := (1 - float64(distance)/float64(maxLen)) * 100
	return similarity
}

// LevenshteinDistance calculates the Levenshtein distance between two strings
func levenshteinDistance(a, b string) int {
	a = ConvertArabicAlphabetToPersian(a)
	b = ConvertArabicAlphabetToPersian(b)

	la, lb := len(a), len(b)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}

	matrix := make([][]int, la+1)
	for i := range matrix {
		matrix[i] = make([]int, lb+1)
	}

	for i := 0; i <= la; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}

			matrix[i][j] = minItem(
				matrix[i-1][j]+1,      // Deletion
				matrix[i][j-1]+1,      // Insertion
				matrix[i-1][j-1]+cost, // Substitution
			)
		}
	}

	return matrix[la][lb]
}

// min returns the minimum of three integers
func minItem(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// max returns the maximum of two integers
func maxItem(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SanitiseSpace(str string) string {
	str = strings.ReplaceAll(str, "\u200C", "") //  remove ZWNJ ( zero-width non-joiner )
	str = strings.TrimSpace(str)
	return str
}

func Sluginfy(input string) string {
	if input == "" {
		return input
	}

	// Convert to lowercase
	slug := strings.ToLower(input)

	// Replace spaces with hyphen
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters using regular expressions
	regex := regexp.MustCompile("[^a-z0-9-]")
	slug = regex.ReplaceAllString(slug, "")

	return slug
}

func EscapeMarkdownV2(text string) string {
	specialChars := `_*[]()~` + "`" + `>#+-=|{}.!`
	for _, char := range specialChars {
		text = strings.ReplaceAll(text, string(char), `\`+string(char))
	}
	return text
}

// CleanFileName returns a sanitized filename safe for most filesystems.
// for example : report<2024>|draft.docx => "report2024draft.docx"
func CleanFileName(fileName string) string {
	// Characters not allowed in Windows file names: < > : " / \ | ? * and control characters
	invalidChars := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)
	clean := invalidChars.ReplaceAllString(fileName, "")

	// Replace multiple spaces with single dash (or underscore if preferred)
	clean = strings.TrimSpace(clean)
	clean = regexp.MustCompile(`\s+`).ReplaceAllString(clean, "-")

	// Optionally, limit length or remove trailing dot
	clean = strings.TrimRight(clean, ".")

	return clean
}

// HashStringHex returns a trimmed hex-encoded SHA-256 hash of the input string.
func HashStringHex(s string, length int) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])[:length]
}
