package ui

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/FYE-0603/hangman/internal/game"
	"github.com/FYE-0603/hangman/internal/wordlist"
)

// RunMenu starts the main application loop with the menu.
func RunMenu(reader *bufio.Reader, provider *wordlist.Provider) {
	for {
		fmt.Println("\n=== MAIN MENU ===")
		fmt.Println("1. Start new game")
		fmt.Println("2. Stats")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			word := provider.Next()
			gameInstance := game.New(word, 6)
			PlayRound(reader, gameInstance)
		case "2":
			fmt.Println("\nStats are not implemented yet. Stay tuned!")
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose 1, 2 or 3.")
		}
	}
}
