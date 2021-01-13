// Package words generates word combinations consisting of an adjective describing an animal.
// It is useful for creating human-readable strings to identify versions of a program (as in the Ubuntu naming scheme).
package words

import (
	"crypto/sha256"
	"math/rand"
	"time"
)

// Get returns pseudorandom words for a given byte slice.
func Get(input []byte) string {
	// hash input data
	hash := sha256.Sum256(input)

	// fold hash to int64
	var seed int64
	for i, e := range hash {
		seed += int64(e) << ((uint(i) % 8) * 8)
	}

	// seed the pseudorandom generator
	rand.Seed(seed)

	// return pseudorandom words
	return adjectives[rand.Intn(len(adjectives))] + " " + animals[rand.Intn(len(animals))]
}

// Random returns almost true random words.
func Random() string {
	// get some randomness
	seed := time.Now().UnixNano()

	// seed the pseudorandom generator
	rand.Seed(seed)

	// return random words
	return adjectives[rand.Intn(len(adjectives))] + " " + animals[rand.Intn(len(animals))]
}

// RandomAdjective returns a random adjective.
// Use the adjective to describe any noun devoid of a description.
func RandomAdjective() string {
	// get some randomness
	seed := time.Now().UnixNano()

	// seed the pseudorandom generator
	rand.Seed(seed)

	// return random words
	return adjectives[rand.Intn(len(adjectives))]
}
