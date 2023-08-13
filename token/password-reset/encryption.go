package passwordreset

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
)

func encryptStruct(data PasswordPayloads, key []byte) (string, error) {
	// Serialize the struct to JSON
	serializedData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Generate a random nonce
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Create the cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt and authenticate the data
	ciphertext := aesGCM.Seal(nil, nonce, serializedData, nil)

	// Combine the nonce and ciphertext, then base64 encode
	encryptedData := append(nonce, ciphertext...)
	encryptedDataString := base64.StdEncoding.EncodeToString(encryptedData)

	return encryptedDataString, nil
}

func decryptStruct(encryptedData string, key []byte) (PasswordPayloads, error) {
	// Base64 decode the encrypted data
	encryptedDataBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return PasswordPayloads{}, err
	}

	// Extract the nonce from the encrypted data
	nonce := encryptedDataBytes[:12]
	ciphertext := encryptedDataBytes[12:]

	// Create the cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return PasswordPayloads{}, err
	}

	// Create GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return PasswordPayloads{}, err
	}

	// Decrypt and verify the data
	serializedData, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return PasswordPayloads{}, err
	}

	// Deserialize the decrypted JSON data into the struct
	var decryptedData PasswordPayloads
	err = json.Unmarshal(serializedData, &decryptedData)
	if err != nil {
		return PasswordPayloads{}, err
	}

	return decryptedData, nil
}
