package main

import (
	"fmt"

	"github.com/carlostrejo2308/GoTakToe/pkg/game"
	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

func main() {
	// Start a new game
	g := game.NewGame()

	// Play the game
	g.Play()

	// Print the winner
	switch g.Winner {
	case piece.Empty:
		fmt.Println("It's a draw!")
	case piece.Human:
		fmt.Println("You win!")
	case piece.Ia:
		fmt.Println("I win!")
	}
}
