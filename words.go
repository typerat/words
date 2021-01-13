// Package words generates word combinations consisting of an adjective describing an animal.
// It is useful for creating human-readable strings to identify versions of a program (as in the Ubuntu naming scheme).
package words

import (
	"time"
)

// Get returns pseudorandom words for a given input.
func Get(input interface{}) string {
	r := rand(input)

	// return pseudorandom words
	return adjectives[r%len(adjectives)] + " " + animals[r%len(animals)]
}

// Random returns almost true random words.
func Random() string {
	// get some randomness
	seed := time.Now().UnixNano()

	// return random words
	return Get(seed)
}

// RandomAdjective returns a random adjective.
// Use the adjective to describe any noun devoid of a description.
func RandomAdjective() string {
	// get some randomness
	seed := time.Now().UnixNano()
	r := rand(seed)

	// return random words
	return adjectives[r%len(adjectives)]
}
