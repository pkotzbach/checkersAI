package main

import (
	"si2/src/si3/checkers"
)

func main() {
	game := checkers.Game{}
	checkers.PopulateNewBoard()
	game.InitPlayers()
	game.GameLoopMiniMax()
	// game.GameLoop()
}
