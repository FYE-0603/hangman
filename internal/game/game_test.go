package game

import "testing"

func TestGuessCorrectLetter(t *testing.T) {
	g := New("go", 6)

	hit, err := g.Guess('g')
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !hit {
		t.Fatal("expected 'g' to be a hit")
	}
	if g.AttemptsLeft() != 6 {
		t.Errorf("attempts should not change on a hit, got %d", g.AttemptsLeft())
	}
}

func TestGuessWrongLetterReducesAttempts(t *testing.T) {
	g := New("go", 6)

	hit, err := g.Guess('z')
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hit {
		t.Fatal("expected 'z' to be a miss")
	}
	if g.AttemptsLeft() != 5 {
		t.Errorf("want 5 attempts left, got %d", g.AttemptsLeft())
	}
}

func TestGuessSameLetterTwiceReturnsError(t *testing.T) {
	g := New("go", 6)
	_, _ = g.Guess('g')

	_, err := g.Guess('g')
	if err != ErrAlreadyGuessed {
		t.Fatalf("want ErrAlreadyGuessed, got %v", err)
	}
}

func TestInvalidLetterReturnsError(t *testing.T) {
	g := New("go", 6)

	if _, err := g.Guess('1'); err != ErrInvalidLetter {
		t.Fatalf("want ErrInvalidLetter, got %v", err)
	}
}

func TestMaskedWord(t *testing.T) {
	g := New("go", 6)
	if got := g.MaskedWord(); got != "_ _" {
		t.Fatalf("want %q, got %q", "_ _", got)
	}

	_, _ = g.Guess('g')
	if got := g.MaskedWord(); got != "g _" {
		t.Fatalf("want %q, got %q", "g _", got)
	}
}

func TestWinCondition(t *testing.T) {
	g := New("go", 6)
	_, _ = g.Guess('g')
	_, _ = g.Guess('o')

	if !g.IsWon() {
		t.Fatal("expected the game to be won")
	}
	if !g.IsOver() {
		t.Fatal("expected the game to be over")
	}
}

func TestLoseCondition(t *testing.T) {
	g := New("go", 2)
	_, _ = g.Guess('x')
	_, _ = g.Guess('y')

	if !g.IsLost() {
		t.Fatal("expected the game to be lost")
	}
	if !g.IsOver() {
		t.Fatal("expected the game to be over")
	}
}
