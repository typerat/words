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
			expected: 3016310759690296599,
		},
		{
			seed:     1,
			expected: 7748076420210162913,
		},
		{
			seed:     "woof",
			expected: 1734376043825659541,
		},
		{
			seed:     []byte("hello world"),
			expected: 4129000111362358792,
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
