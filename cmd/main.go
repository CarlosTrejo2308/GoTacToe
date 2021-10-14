package main

import (
	"fmt"

	"github.com/carlostrejo2308/GoTakToe/pkg/game"
	"github.com/carlostrejo2308/GoTakToe/pkg/ia"
	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

func main() {
	// Start a new game
	g := game.NewGame()
	g.SetIa(ia.Challenger)

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
