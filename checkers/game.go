package checkers

import (
	"fmt"
	"time"
)

const WIN = 10000
const LOSE = -WIN

type Game struct {
	players [2]Player
}

func (g *Game) InitPlayers() {
	g.players[0] = Player{Number: 1, Symbol: "X", PieceVector: -1, KingMoves: 0}
	g.players[1] = Player{Number: 2, Symbol: "O", PieceVector: 1, KingMoves: 0}
}

func (g *Game) GameLoop() {
	for {
		if g.playing() {
			g.players[0].PlayerTurn()
		} else {
			fmt.Println("player 2 won")
			fmt.Println(g.GameState())
			break
		}
		if g.playing() {
			g.players[1].PlayerTurn()
		} else {
			fmt.Println("player 1 won")
			fmt.Println(g.GameState())
			break
		}
	}
	fmt.Println("END")
}

func (g *Game) GameLoopMiniMax() {
	for {
		if g.playing() {
			fmt.Println(BoardAsString())
			start := time.Now()

			_, move, checked := g.MinMax(6, 0)
			g.players[0].playerTurnLogic(move, g.players[0].getValidMovesWithMultipleCapture())
			fmt.Println("bot move", boardToPos(move.from), "->", boardToPos(move.to))
			fmt.Println("move took him", time.Since(start))
			fmt.Println("checked", checked, "nodes")
		} else {
			fmt.Println("player 2 won")
			fmt.Println(g.GameState())
			break
		}
		if g.playing() {
			fmt.Println()
			g.players[1].PlayerTurn()
		} else {
			fmt.Println("player 1 won")
			fmt.Println(g.GameState())
			break
		}
	}
	fmt.Println("END")
}

func (g *Game) isDraw() bool {
	// fmt.Println(g.players[0].KingMoves, g.players[1].KingMoves)
	return g.players[0].KingMoves >= 15 || g.players[1].KingMoves >= 15
}

func (g *Game) playing() bool {
	gs := g.GameState()
	// fmt.Println(gs != WIN, !g.isDraw(), gs != LOSE)
	return gs != WIN && !g.isDraw() && gs != LOSE
}

func (g *Game) GameState() int {
	if !g.players[1].hasValidMoves() {
		return WIN
	}
	if !g.players[0].hasValidMoves() {
		return LOSE

	}
	return g.players[0].calculateCost() - g.players[1].calculateCost()
}

//player0 is player, player1 is enemy
func (g *Game) MinMax(depth, currentPlayer int) (int, Move, int) {
	if state := g.GameState(); depth == 0 || state == WIN || state == LOSE || g.isDraw() {
		return state, Move{[2]int{-1, -1}, [2]int{-1, -1}, make([][2]int, 0)}, 0
	}
	moves := g.players[currentPlayer].getValidMovesWithMultipleCapture()
	min := WIN
	max := LOSE
	val := 0
	var checked int
	checkedAll := 1
	var (
		minMove Move
		maxMove Move
	)

	for _, move := range moves {
		backupBoard := board
		backupKingMoves := g.players[currentPlayer].KingMoves

		g.players[currentPlayer].playerTurnLogic(move, moves)
		val, _, checked = g.MinMax(depth-1, (currentPlayer+1)%2)

		if val < min {
			min = val
			minMove = move
		}
		if val > max {
			max = val
			maxMove = move
		}

		board = backupBoard
		g.players[currentPlayer].KingMoves = backupKingMoves
		checkedAll += checked
	}
	if currentPlayer == 0 {
		// fmt.Println(depth, max, boardToPos(maxMove.from), "->", boardToPos(maxMove.to))
		return max, maxMove, checkedAll
	}
	// fmt.Println(depth, min, boardToPos(minMove.from), "->", boardToPos(minMove.to))
	return min, minMove, checkedAll
}
