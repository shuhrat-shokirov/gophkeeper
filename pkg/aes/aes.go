package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"go.uber.org/fx"

	"gophkeeper/pkg/config"
)

// secretKey is the secret key used for encryption and decryption
// loaded while application starts
var (
	secretKey string

	Module = fx.Invoke(New)
)

type Params struct {
	fx.In

	Config config.Config
}

func New(params Params) {
	secretKey = params.Config.GetString("aes.secret_key")
}

func MustEncrypt(plainText string) string {
	cipherText, err := encryptB64(plainText)
	if err != nil {
		return plainText
	}

	return cipherText
}

func MustDecrypt(b64CipherText string) string {
	plainText, err := decryptB64(b64CipherText)
	if err != nil {
		return b64CipherText
	}

	return plainText
}

// encryptB64 encrypts the plainText using AES-128 encryption algorithm
// and returns base64 encoded string.
// First block is randomly generated iv and the rest is cipher text
func encryptB64(plainText string) (string, error) {
	cipherText, err := encrypt([]byte(secretKey), []byte(plainText))
	if err != nil {
		return "", fmt.Errorf("encrypt error: failed to encrypt: %w", err)
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// decryptB64 decrypts the base64 encoded string using AES-128 decryption algorithm
// and returns plain text.
// Reads iv from first block of cipherText else returns error
func decryptB64(b64CipherText string) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(b64CipherText)
	if err != nil {
		return "", fmt.Errorf("decrypt error: failed to decode base64: %w", err)
	}

	plainText, err := decrypt([]byte(secretKey), encryptedData)
	if err != nil {
		return "", fmt.Errorf("decrypt error: failed to decrypt: %w", err)
	}

	return string(plainText), nil
}

func encrypt(key, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("encrypt error: failed to create new cipher: %w", err)
	}

	var (
		cipherText = make([]byte, block.BlockSize()+len(plainText))
		iv         = cipherText[:block.BlockSize()]
	)

	// read random iv from rand.Reader
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("encrypt error: failed to read random iv: %w", err)
	}

	// create new CFB encrypter stream using block and iv
	stream := cipher.NewCFBEncrypter(block, iv)
	// encrypt the plainText and write to cipherText
	stream.XORKeyStream(cipherText[block.BlockSize():], plainText)

	return cipherText, nil
}

func decrypt(key, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("decrypt error: failed to create new cipher: %w", err)
	}

	if len(cipherText) < block.BlockSize() {
		return nil, fmt.Errorf("decrypt error: cipherText is less than minimum block size")
	}

	var (
		iv     = cipherText[:block.BlockSize()]
		stream = cipher.NewCFBDecrypter(block, iv)
	)
	cipherText = cipherText[block.BlockSize():]

	// decrypt the cipherText and write to cipherText
	// to avoid usage of extra memory
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}
