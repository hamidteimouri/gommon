package gommon

import (
	"crypto/rand"
	"github.com/google/uuid"
	"math/big"
	mathRand "math/rand"
	"strings"
	"unsafe"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString deprecated
func RandomString(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

const charsets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomStringCode(length int) string {
	sb := make([]byte, length)
	for i := range sb {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charsets))))
		if err != nil {
			return RandomString(length)
		}
		sb[i] = charsets[num.Int64()]
	}
	return string(sb)
}

func RandomOtpCode(length int) string {
	// Define the character set from which the code will be generated
	charset := "0123456789" // You can include other characters if needed

	// Generate the verification code
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[mathRand.Intn(len(charset))]
	}

	return string(code)
}

// GenerateUniqueCode generates a unique lowercase code using a prefix and count of characters.
// The code includes `count` characters randomly picked from different parts of the UUID.
func GenerateUniqueCode(prefix string, count int) string {
	u := uuid.New().String()
	u = strings.ReplaceAll(u, "-", "") // remove dashes to simplify indexing

	var builder strings.Builder
	builder.WriteString(prefix)

	// Spread out the characters by jumping evenly through the UUID
	step := len(u) / count
	for i := 0; i < count; i++ {
		builder.WriteByte(u[i*step])
	}

	return strings.ToLower(builder.String())
}
