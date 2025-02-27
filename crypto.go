package gommon

import (
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
