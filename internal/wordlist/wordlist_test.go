package wordlist

import (
	"testing"
)

func TestNextDoesNotRepeatUntilExhausted(t *testing.T) {
	// Create a provider with exactly 3 words.
	words := []string{"a", "b", "c"}
	p := NewProvider(words)

	// We'll collect the first 3 words it gives us.
	seen := make(map[string]bool)
	for i := 0; i < 3; i++ {
		w := p.Next()
		if seen[w] {
			t.Fatalf("word %q was returned twice before the pool was exhausted", w)
		}
		seen[w] = true
	}
}
