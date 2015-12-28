package words

import (
	"crypto/sha256"
	"math/rand"
	"time"
)

// get pseudorandom words for given byte slice
func Get(input []byte) string {
	// hash input data
	hash := sha256.Sum256(input)

	// fold hash to int64
	var seed int64
	for i, e := range hash {
		seed += int64(e) << ((uint(i) % 8) * 8)
	}

	// set seed for pseudorandom generator
	rand.Seed(seed)

	// return pseudorandom words
	return Adjectives[rand.Intn(len(Adjectives))] + " " + Animals[rand.Intn(len(Animals))]
}

// get almost true random words
func Random() string {
	// get some randomness
	b := time.Now().UnixNano()

	// seed the pseudorandom generator
	rand.Seed(b)

	// return random words
	return Adjectives[rand.Intn(len(Adjectives))] + " " + Animals[rand.Intn(len(Animals))]
}
