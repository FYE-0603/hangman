package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FYE-0603/hangman/internal/ui"
	"github.com/FYE-0603/hangman/internal/wordlist"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	provider, err := wordlist.LoadFromFile()
	if err != nil {
		fmt.Println("Error loading wordlist:", err)
		os.Exit(1)
	}

	ui.RunMenu(reader, provider)
}
