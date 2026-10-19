package ui

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/FYE-0603/hangman/internal/game"
	"github.com/FYE-0603/hangman/internal/history"
	"github.com/FYE-0603/hangman/internal/wordlist"
)

// RunMenu starts the main application loop with the menu.
func RunMenu(reader *bufio.Reader, provider *wordlist.Provider) {
	stats, err := history.Load()
	if err != nil {
		fmt.Println("Warning: could not load stats:", err)
		stats = &history.Stats{}
	}

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
			won := PlayRound(reader, gameInstance)

			stats.TotalGames++
			if won {
				stats.Wins++
			} else {
				stats.Losses++
			}
			if err := stats.Save(); err != nil {
				fmt.Println("Warning: could not save stats:", err)
			}
		case "2":
			fmt.Printf("\n--- STATS ---\n")
			fmt.Printf("Total games: %d\n", stats.TotalGames)
			fmt.Printf("Wins:        %d\n", stats.Wins)
			fmt.Printf("Losses:      %d\n", stats.Losses)
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose 1, 2 or 3.")
		}
	}
}
