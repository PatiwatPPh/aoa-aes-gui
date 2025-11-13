package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	keySize = 32 // AES-256 requires 32 bytes key
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GenerateKey generates a new AES-256 key
func (a *App) GenerateKey() string {
	key := make([]byte, keySize)
	if _, err := rand.Read(key); err != nil {
		return ""
	}
	return hex.EncodeToString(key)
}

// Encrypt encrypts the plaintext using the provided key
func (a *App) Encrypt(keyInput, plaintext string) (string, error) {
	// Parse key
	key, err := parseKey(keyInput)
	if err != nil {
		return "", err
	}

	// Encrypt
	encrypted, err := encrypt([]byte(plaintext), key)
	if err != nil {
		return "", err
	}

	// Encode to base64
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts the encrypted text using the provided key
func (a *App) Decrypt(keyInput, encryptedBase64 string) (string, error) {
	// Parse key
	key, err := parseKey(keyInput)
	if err != nil {
		return "", err
	}

	// Decode base64
	encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", fmt.Errorf("invalid base64 format: %v", err)
	}

	// Decrypt
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// parseKey parses key from either string format (32 bytes) or hex format (64 characters)
func parseKey(keyInput string) ([]byte, error) {
	// Try as direct string first (32 bytes)
	if len(keyInput) == keySize {
		return []byte(keyInput), nil
	}

	// Try as hex format (64 characters)
	if len(keyInput) == keySize*2 {
		key, err := hex.DecodeString(keyInput)
		if err != nil {
			return nil, fmt.Errorf("invalid hex key format: %v", err)
		}
		return key, nil
	}

	return nil, fmt.Errorf("key must be either %d bytes (string) or %d hex characters, got %d characters", keySize, keySize*2, len(keyInput))
}

// encrypt encrypts plaintext using AES-256-GCM
func encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt and append nonce at the beginning
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decrypt decrypts ciphertext using AES-256-GCM
func decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Extract nonce and ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
