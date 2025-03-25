package gommon

import (
	"crypto/rand"
	"math/big"
	mathRand "math/rand"
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
