package words

import (
	"crypto/sha256"
	"math/rand"
	"time"
)

func Get(input []byte) string {
	hash := sha256.Sum256(input)

	var seed int64
	for i, e := range hash {
		seed += int64(e) << ((uint(i) % 8) * 8)
	}
	rand.Seed(seed)
	return Adjectives[rand.Intn(len(Adjectives))] + " " + Animals[rand.Intn(len(Animals))]
}

func Random() string {
	b := time.Now().UnixNano()

	rand.Seed(b)

	return Adjectives[rand.Intn(len(Adjectives))] + " " + Animals[rand.Intn(len(Animals))]
}
