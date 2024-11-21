package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

// Encrypt string message using AES encryption. Make sure AES_SECRET
// key with value was added to environmental variables.
func EncryptAES(message string) (string, error) {
	var (
		key        = os.Getenv("AES_SECRET")
		plainText  = []byte(message)
		block, err = aes.NewCipher([]byte(key))
	)

	if err != nil {
		return "", err
	}

	var (
		cipherText = make([]byte, aes.BlockSize+len(plainText))
		iv         = cipherText[:aes.BlockSize]
	)

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText), nil
}

// Decrypt encrypted message using AES encryption. Make sure AES_SECRET
// key with value was added to environmental variables.
func DecryptAES(secure string) (decoded string, err error) {
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)
	if err != nil {
		return
	}

	key := os.Getenv("AES_SECRET")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}
