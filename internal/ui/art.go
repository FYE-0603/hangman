package ui

// hangmanStages contains the ASCII art for each number of remaining attempts.
// Index 0 is the "almost lost" state, index 6 is the "clean gallows" state.
var hangmanStages = []string{
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
}

// HangmanArt returns the ASCII art for the given number of remaining attempts.
func HangmanArt(attemptsLeft int) string {
	if attemptsLeft < 0 {
		attemptsLeft = 0
	}
	if attemptsLeft >= len(hangmanStages) {
		attemptsLeft = len(hangmanStages) - 1
	}
	return hangmanStages[attemptsLeft]
}
