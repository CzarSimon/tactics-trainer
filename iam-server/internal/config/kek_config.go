package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
)

// LoadKEKConfig reads and parses key encryption keys from a given filepath
func LoadKEKConfig(path string) []models.KeyEncryptionKey {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		log.Panic("failed to read key encryption keys. Error: %w", err)
	}

	lines := strings.Split(string(bytes), "\n")
	keys := make([]models.KeyEncryptionKey, 0)
	for i, line := range lines {
		kek, err := models.ParseKeyEncryptionKey(line)
		if err != nil {
			log.Panic("failed parse key encryption key on line %d. Error: %w", i+1, err)
		}

		keys = append(keys, kek)
	}

	return keys
}
