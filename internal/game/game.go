package game

import (
	"sort"
	"strings"
	"unicode"
)

// Game holds the state of a single hangman round.
type Game struct {
	word         string
	guessed      map[rune]bool
	attemptsLeft int
	maxAttempts  int
}

// New creates a new game for the given word and attempt limit.
func New(word string, maxAttempts int) *Game {
	return &Game{
		word:         strings.ToLower(word),
		guessed:      make(map[rune]bool),
		attemptsLeft: maxAttempts,
		maxAttempts:  maxAttempts,
	}
}

// Guess processes a single letter guess.
// It returns true if the letter is in the word.
func (g *Game) Guess(letter rune) (bool, error) {
	if g.IsOver() {
		return false, ErrGameOver
	}

	letter = unicode.ToLower(letter)

	if letter < 'a' || letter > 'z' {
		return false, ErrInvalidLetter
	}

	if g.guessed[letter] {
		return false, ErrAlreadyGuessed
	}

	g.guessed[letter] = true

	hit := strings.ContainsRune(g.word, letter)
	if !hit {
		g.attemptsLeft--
	}

	return hit, nil
}

// MaskedWord returns the word with unguessed letters replaced by underscores.
// Non-letter characters (spaces, hyphens) are always shown as is.
func (g *Game) MaskedWord() string {
	parts := make([]string, 0, len(g.word))

	for _, r := range g.word {
		switch {
		case r == ' ' || r == '-':
			parts = append(parts, string(r))
		case g.guessed[r]:
			parts = append(parts, string(r))
		default:
			parts = append(parts, "_")
		}
	}

	return strings.Join(parts, " ")
}

// IsWon reports whether every letter of the word has been guessed.
func (g *Game) IsWon() bool {
	for _, r := range g.word {
		if !g.guessed[r] {
			return false
		}
	}

	return true
}

// IsLost reports whether the player ran out of attempts.
func (g *Game) IsLost() bool {
	return g.attemptsLeft <= 0 && !g.IsWon()
}

// IsOver reports whether the game has finished (win or lose).
func (g *Game) IsOver() bool {
	return g.IsWon() || g.IsLost()
}

// AttemptsLeft returns the number of remaining attempts.
func (g *Game) AttemptsLeft() int { return g.attemptsLeft }

// MaxAttempts returns the initial number of attempts.
func (g *Game) MaxAttempts() int { return g.maxAttempts }

// Word returns the hidden word. Only use this after the game is over.
func (g *Game) Word() string { return g.word }

// GuessedLetters returns all guessed letters, sorted alphabetically.
func (g *Game) GuessedLetters() []string {
	letters := make([]string, 0, len(g.guessed))
	for r := range g.guessed {
		letters = append(letters, string(r))
	}
	sort.Strings(letters)
	return letters
}
