package words

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleGet() {
	words := Get("some seed")
	fmt.Println(words)
	// Output: flirtatious tarantula
}

// random tests are quite useless, but it feels good to have 100% coverage

func TestRandom(t *testing.T) {
	words := Random()

	parts := strings.Split(words, " ")

	if len(parts) != 2 {
		t.Error("expected two words")
	}
}

func TestRandomAdjective(t *testing.T) {
	word := RandomAdjective()

	if len(word) == 0 {
		t.Error("expected a word")
	}
}
