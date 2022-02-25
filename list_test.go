package words

import (
	"strings"
	"testing"
)

func TestAdjectives(t *testing.T) {
	exists := map[string]bool{}

	for _, word := range adjectives {
		trimmed := strings.TrimSpace(word)
		if trimmed != word {
			t.Errorf("adjective has whitespace: %s", word)
		}
		if exists[word] {
			t.Errorf("duplicate adjective: %s", word)
		}

		exists[word] = true
	}
}

func TestAnimals(t *testing.T) {
	exists := map[string]bool{}

	for _, word := range animals {
		trimmed := strings.TrimSpace(word)
		if trimmed != word {
			t.Errorf("animal has whitespace: %s", word)
		}

		if exists[word] {
			t.Errorf("duplicate animal: %s", word)
		}

		exists[word] = true
	}
}
