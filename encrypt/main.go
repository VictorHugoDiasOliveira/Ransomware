package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	defer timer()()

	// Read file with key
	key, _ := os.ReadFile("./key.txt")

	// Starting a block with Cipher
	block, _ := aes.NewCipher(key)

	// Starting GCM
	gcm, _ := cipher.NewGCM(block)

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat
	nonce := make([]byte, gcm.NonceSize())
	// if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	// 	log.Fatalf("Nonce error: %v", err.Error())
	// }

	// Walk through directory and subdirectory
	filepath.Walk("./test/", filepath.WalkFunc(func(path string, file os.FileInfo, err error) error {

		// Do not interact with Directory
		if !file.IsDir() {

			// Read content to encrypt
			plainText, _ := os.ReadFile(path)

			// Encrypt plainText
			cipherText := gcm.Seal(nonce, nonce, plainText, nil)

			// Write encrypted content
			os.WriteFile(path, cipherText, 0777)
		}
		return nil
	}))
}

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("It took: %v", time.Since(start))
	}
}
