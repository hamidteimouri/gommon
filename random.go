package gommon

import (
	"crypto/rand"
	mathRand "math/rand"
	"unsafe"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
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
