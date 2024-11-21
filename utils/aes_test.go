package utils

import (
	"encoding/base64"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptAES(t *testing.T) {
	// Set the AES_SECRET environment variable for testing
	require.NoError(t, os.Setenv("AES_SECRET", "thisis16byteskey"), "Failed to set AES_SECRET")

	tests := []struct {
		name       string
		input      string
		shouldFail bool
	}{
		{"Basic string", "Hello, World!", false},
		{"Empty string", "", false},
		{"Special characters", "!@#$%^&*()", false},
		{"Unicode characters", "你好，世界🌟", false},
		{"Long string", string(make([]byte, 1000)), false},
		{"Tampered ciphertext", "tampered==ciphertext", true},
		{"Invalid Base64 input", "invalid-base64", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldFail {
				// Test decryption failure cases
				_, err := DecryptAES(tt.input)
				assert.Error(t, err, "Expected decryption to fail for input: %s", tt.input)
			} else {
				// Test encryption
				encrypted, err := EncryptAES(tt.input)
				require.NoError(t, err, "Encryption failed for input: %s", tt.input)

				// Test decryption
				decrypted, err := DecryptAES(encrypted)
				require.NoError(t, err, "Decryption failed for encrypted message")
				assert.Equal(t, tt.input, decrypted, "Decrypted message does not match original")
			}
		})
	}
}

func TestEnvironmentVariable(t *testing.T) {
	// Unset the AES_SECRET environment variable
	require.NoError(t, os.Unsetenv("AES_SECRET"), "Failed to unset AES_SECRET")

	_, err := EncryptAES("test")
	assert.Error(t, err, "Expected encryption to fail due to missing AES_SECRET")

	_, err = DecryptAES("test")
	assert.Error(t, err, "Expected decryption to fail due to missing AES_SECRET")

	// Reset the environment variable
	require.NoError(t, os.Setenv("AES_SECRET", "thisis16byteskey"), "Failed to reset AES_SECRET")
}

func TestShortCiphertext(t *testing.T) {
	// Test decryption with a short ciphertext
	shortCiphertext := base64.RawStdEncoding.EncodeToString([]byte("short"))
	_, err := DecryptAES(shortCiphertext)
	assert.Error(t, err, "Expected failure for short ciphertext")
}

func TestKeyLengthValidity(t *testing.T) {
	tests := []struct {
		name       string
		aesSecret  string
		shouldFail bool
	}{
		{"Invalid key length - too short", "shortkey", true},
		{"Invalid key length - too long", "thiskeyiswaytoolongforaes!", true},
		{"Valid key length - 16 bytes", "thisis16byteskey", false},
		{"Valid key length - 24 bytes", "thisis24byteslongkey1234", false},
		{"Valid key length - 32 bytes", "thisis32byteslongkeyfortesting!!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("AES_SECRET", tt.aesSecret)

			_, encryptErr := EncryptAES("test message")
			if tt.shouldFail {
				assert.Error(t, encryptErr, "Expected encryption to fail for key: %s", tt.aesSecret)
			} else {
				assert.NoError(t, encryptErr, "Expected encryption to succeed for key: %s", tt.aesSecret)
			}
		})
	}
}

func TestRandomInitializationVector(t *testing.T) {
	require.NoError(t, os.Setenv("AES_SECRET", "thisis16byteskey"), "Failed to set AES_SECRET")

	message := "test message"
	var ciphertexts []string

	for i := 0; i < 5; i++ {
		encrypted, err := EncryptAES(message)
		require.NoError(t, err, "Encryption failed")
		ciphertexts = append(ciphertexts, encrypted)
	}

	for i := 0; i < len(ciphertexts)-1; i++ {
		assert.NotEqual(t, ciphertexts[i], ciphertexts[i+1], "Ciphertexts should differ due to random IVs")
	}
}

func TestKeySensitivity(t *testing.T) {
	require.NoError(t, os.Setenv("AES_SECRET", "thisis16byteskey"), "Failed to set AES_SECRET")

	message := "sensitive data"
	encrypted, err := EncryptAES(message)
	require.NoError(t, err, "Encryption failed")

	os.Setenv("AES_SECRET", "wrong16byteskey")
	_, decryptErr := DecryptAES(encrypted)
	assert.Error(t, decryptErr, "Expected decryption to fail with incorrect key")
}

func TestStress(t *testing.T) {
	require.NoError(t, os.Setenv("AES_SECRET", "thisis16byteskey"), "Failed to set AES_SECRET")

	message := "stress test message"
	for i := 0; i < 1000; i++ {
		t.Run("Iteration "+strconv.Itoa(i), func(t *testing.T) {
			encrypted, err := EncryptAES(message)
			require.NoError(t, err, "Encryption failed")

			decrypted, err := DecryptAES(encrypted)
			require.NoError(t, err, "Decryption failed")
			assert.Equal(t, message, decrypted, "Decrypted message does not match original")
		})
	}
}

func TestConcurrency(t *testing.T) {
	require.NoError(t, os.Setenv("AES_SECRET", "thisis16byteskey"), "Failed to set AES_SECRET")

	message := "concurrent test message"
	const goroutines = 50

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			encrypted, err := EncryptAES(message)
			assert.NoError(t, err, "Encryption failed")

			decrypted, err := DecryptAES(encrypted)
			assert.NoError(t, err, "Decryption failed")
			assert.Equal(t, message, decrypted, "Decrypted message does not match original")
		}()
	}
	wg.Wait()
}

func TestCrossSessionConsistency(t *testing.T) {
	require.NoError(t, os.Setenv("AES_SECRET", "thisis16byteskey"), "Failed to set AES_SECRET")

	message := "cross session message"
	encrypted, err := EncryptAES(message)
	require.NoError(t, err, "Encryption failed")

	// Simulating a new session
	os.Setenv("AES_SECRET", "thisis16byteskey")
	decrypted, err := DecryptAES(encrypted)
	require.NoError(t, err, "Decryption failed in new session")
	assert.Equal(t, message, decrypted, "Decrypted message does not match original")
}
