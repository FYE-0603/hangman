package game

import "errors"

var (
	// ErrInvalidLetter is returned when the input is not a latin letter a-z.
	ErrInvalidLetter = errors.New("letter must be a single latin character a-z")

	// ErrAlreadyGuessed is returned when the same letter is guessed twice.
	ErrAlreadyGuessed = errors.New("letter was already guessed")

	// ErrGameOver is returned when a guess is made after the game has ended.
	ErrGameOver = errors.New("game is already over")
)
