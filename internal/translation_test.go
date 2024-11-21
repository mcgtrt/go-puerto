package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslationManager(t *testing.T) {
	tm := NewTranslationManager()

	// Test: NewTranslationManager initializes correctly
	t.Run("NewTranslationManager initializes correctly", func(t *testing.T) {
		assert.NotNil(t, tm)
		assert.NotNil(t, tm.translations)
		assert.Equal(t, 0, len(tm.translations), "Expected empty translations map")
	})

	// Test: Load valid translation file
	t.Run("Load valid translation file", func(t *testing.T) {
		lang := "en"
		content := `{"greeting": "Hello", "farewell": "Goodbye"}`
		filePath := "test_translations_en.json"

		// Create a temporary test file
		err := os.WriteFile(filePath, []byte(content), 0644)
		assert.NoError(t, err, "Failed to create test file")
		defer os.Remove(filePath)

		err = tm.Load(lang, filePath)
		assert.NoError(t, err, "Failed to load valid translation file")
		assert.Equal(t, "Hello", tm.Translate(lang, "greeting"))
		assert.Equal(t, "Goodbye", tm.Translate(lang, "farewell"))
	})

	// Test: Load invalid JSON file
	t.Run("Load invalid JSON file", func(t *testing.T) {
		lang := "en"
		content := `{"greeting": "Hello", "farewell": "Goodbye"` // Missing closing brace
		filePath := "test_invalid_translations.json"

		// Create a temporary test file
		err := os.WriteFile(filePath, []byte(content), 0644)
		assert.NoError(t, err, "Failed to create test file")
		defer os.Remove(filePath)

		err = tm.Load(lang, filePath)
		assert.Error(t, err, "Expected error for invalid JSON")
	})

	// Test: Load non-existent file
	t.Run("Load non-existent file", func(t *testing.T) {
		lang := "en"
		filePath := "non_existent_file.json"

		err := tm.Load(lang, filePath)
		assert.Error(t, err, "Expected error for non-existent file")
	})

	// Test: Translate with missing key
	t.Run("Translate with missing key", func(t *testing.T) {
		lang := "en"
		assert.Equal(t, "", tm.Translate(lang, "missing_key"), "Expected empty string for missing key")
	})

	// Test: Translate with missing language
	t.Run("Translate with missing language", func(t *testing.T) {
		assert.Equal(t, "", tm.Translate("fr", "greeting"), "Expected empty string for missing language")
	})
}
