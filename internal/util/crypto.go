package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// Encrypt encrypts plain text string into cipher text string using a key.
func Encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(createHash(key)[:32]))
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts cipher text string into plain text string using a key.
func Decrypt(key, cryptoText string) (string, error) {
	block, err := aes.NewCipher([]byte(createHash(key)[:32]))
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// createHash creates a hash from the key for AES.
func createHash(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}
