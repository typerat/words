package words

import (
	"fmt"
	"testing"
)

func TestRandomInt(t *testing.T) {
	testCases := []struct {
		seed     interface{}
		expected int
	}{
		{
			seed:     nil,
			expected: 702289575,
		},
		{
			seed:     1,
			expected: 1803989619,
		},
		{
			seed:     "woof",
			expected: 403815890,
		},
		{
			seed:     []byte("hello world"),
			expected: 961357753,
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.seed), func(t *testing.T) {
			i := rand(tC.seed)
			if i != tC.expected {
				t.Errorf("want: %d, got: %d", tC.expected, i)
			}
		})
	}
}
