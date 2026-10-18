package wordlist

import (
	"bufio"
	"embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed words.txt
var embeddedWords embed.FS

// Provider supplies random words for the game.
type Provider struct {
	words     []string
	remaining []string
}

// NewProvider creates a provider from the given list of words.
func NewProvider(words []string) *Provider {
	remaining := make([]string, len(words))
	copy(remaining, words)

	return &Provider{
		words:     words,
		remaining: remaining,
	}
}

// Random returns a random word from the provider's list.
func (p *Provider) Random() string {
	return p.words[rand.Intn(len(p.words))]
}

// Next returns a random word that has not been returned yet.
// When all words are exhausted, the pool is reset and reshuffled.
func (p *Provider) Next() string {
	if len(p.remaining) == 0 {
		p.remaining = make([]string, len(p.words))
		copy(p.remaining, p.words)
	}

	i := rand.Intn(len(p.remaining))
	word := p.remaining[i]

	// Remove the word at index i by swapping it with the last element,
	// then shrinking the slice by one.
	p.remaining[i] = p.remaining[len(p.remaining)-1]
	p.remaining = p.remaining[:len(p.remaining)-1]

	return word
}

// LoadFromFile reads words from the embedded file, one word per line.
func LoadFromFile() (*Provider, error) {
	file, err := embeddedWords.Open("words.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		words = append(words, strings.ToLower(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(words) == 0 {
		return nil, fmt.Errorf("embedded file contains no valid words")
	}

	return NewProvider(words), nil
}
