package main

import (
	"fmt"

	"github.com/carlostrejo2308/GoTakToe/pkg/game"
	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

func main() {
	g := game.NewGame()
	g.Play()

	switch g.Winner {
	case piece.Empty:
		fmt.Println("It's a draw!")
	case piece.Human:
		fmt.Println("You win!")
	case piece.Ia:
		fmt.Println("I win!")
	}
}
