package internal

import (
	"encoding/json"
	"os"
	"sync"
)

type TranslationManager struct {
	translations map[string]map[string]string // language -> key -> value
	mu           sync.RWMutex
}

func NewTranslationManager() *TranslationManager {
	return &TranslationManager{
		translations: make(map[string]map[string]string),
	}
}

func (tm *TranslationManager) Load(lang, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var translations map[string]string
	if err := json.Unmarshal(data, &translations); err != nil {
		return err
	}

	tm.mu.Lock()
	tm.translations[lang] = translations
	tm.mu.Unlock()
	return nil
}

func (tm *TranslationManager) Translate(lang, key string) string {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.translations[lang][key]
}
