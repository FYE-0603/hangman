package ui

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/FYE-0603/hangman/internal/game"
)

// PlayRound runs a single game of hangman from start to finish.
// It returns true if the player won, false otherwise.
func PlayRound(reader *bufio.Reader, g *game.Game) bool {
	fmt.Println("=== HANGMAN ===")

	for !g.IsOver() {
		drawState(g)

		letter, err := readLetter(reader)
		if err != nil {
			fmt.Printf("  ! %v\n\n", err)
			continue
		}

		hit, err := g.Guess(letter)
		if err != nil {
			fmt.Printf("  ! %v\n\n", err)
			continue
		}

		if hit {
			fmt.Printf("  + '%c' is not the word!\n\n", letter)
		} else {
			fmt.Printf("  - '%c' in not in the word. Attempts left: %d\n\n", letter, g.AttemptsLeft())
		}
	}

	drawState(g)
	if g.IsWon() {
		fmt.Printf("  You won! The word was %q. \n\n", g.Word())
		return true
	}
	fmt.Printf("  You lost. The word was %q. \n\n", g.Word())
	return false
}

// readLetter prompts the user until a single latin letter is entered.
func readLetter(reader *bufio.Reader) (rune, error) {
	fmt.Print("Guess a letter: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	runes := []rune(input)
	if len(runes) != 1 {
		return 0, fmt.Errorf("please enter exactly one letter")
	}

	return runes[0], nil
}

// drawState prints the current hangman picture, masked word, and attempts left.
func drawState(g *game.Game) {
	fmt.Println(HangmanArt(g.AttemptsLeft()))
	fmt.Printf("  Word: %s\n", g.MaskedWord())
	fmt.Printf("  Attempts left: %d/%d\n\n", g.AttemptsLeft(), g.MaxAttempts())
	fmt.Printf("  Guessed: %s\n\n", g.GuessedLetters())
}
