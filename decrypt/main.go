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

	// Reading key
	key, _ := os.ReadFile("./key.txt")

	// Creating block of algorithm
	block, _ := aes.NewCipher(key)

	// Creating GCM mode
	gcm, _ := cipher.NewGCM(block)

	// Walk through directory and subdirectory
	filepath.Walk("./test/", filepath.WalkFunc(func(path string, file os.FileInfo, err error) error {

		// Do not interact with Directory
		if !file.IsDir() {

			// Reading encrypted file content
			cipherText, _ := os.ReadFile(path)

			// Deattached nonce and decrypt
			nonce := cipherText[:gcm.NonceSize()]
			cipherText = cipherText[gcm.NonceSize():]
			plainText, _ := gcm.Open(nil, nonce, cipherText, nil)

			// Writing decrypted content
			os.WriteFile(path, plainText, 0777)
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
