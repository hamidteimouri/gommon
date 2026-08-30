package gommon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func Base64ToHex(base64str string) (string, error) {
	// Decode the Base64 string into binary data
	binary, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(binary), nil
}

func HexToBase64(hexStr string) (string, error) {
	// Decode the hex string into binary data
	binary, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}

	// Encode the binary data into a Base64 string
	return base64.StdEncoding.EncodeToString(binary), nil
}

func HmacSign(secret string, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

func HmacVerify(secret string, data string, signature string) bool {
	expected := HmacSign(secret, data)

	return hmac.Equal(
		[]byte(expected),
		[]byte(signature),
	)
}
