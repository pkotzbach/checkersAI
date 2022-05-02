package checkers

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

type Player struct {
	Number      int
	PieceVector int
	Symbol      string
	KingMoves   int
}

type Move struct {
	from    [2]int
	to      [2]int
	capture [2]int
}

func (m *Move) calcDistance() int {
	return int(math.Abs(float64(m.to[0] - m.from[0])))
}

func (p *Player) getMovesToCheckForPiece(i, j int) []Move {
	var moves []Move

	if isPlayersKing(i, j, p.Number) {
		canKingMove := func(i, j int, enemies *int) bool {
			if onBoard(i, j) {
				if *enemies > 1 {
					return false
				}
				if belongsToEnemy(i, j, p.Number) {
					*enemies++
				}
				return true
			}
			return false
		}

		enemies1 := 0
		enemies2 := 0
		enemies3 := 0
		enemies4 := 0
		border := boardSize
		for k := 1; k <= border; k++ {
			if canKingMove(i+k, j+k, &enemies1) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + k, j + k}, [2]int{-1, -1}})
			}
			if canKingMove(i+k, j-k, &enemies2) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + k, j - k}, [2]int{-1, -1}})
			}
			if canKingMove(i-k, j+k, &enemies3) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i - k, j + k}, [2]int{-1, -1}})
			}
			if canKingMove(i-k, j-k, &enemies4) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i - k, j - k}, [2]int{-1, -1}})
			}
		}
	} else if isPlayersPawn(i, j, p.Number) {
		border := 2
		for k := 1; k <= border; k++ {
			if onBoard(i+(k*p.PieceVector), j+k) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + (k * p.PieceVector), j + k}, [2]int{-1, -1}})
			}
			if onBoard(i+(k*p.PieceVector), j-k) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + (k * p.PieceVector), j - k}, [2]int{-1, -1}})
			}
		}
		if onBoard(i-(border*p.PieceVector), j+border) {
			moves = append(moves, Move{[2]int{i, j}, [2]int{i - (border * p.PieceVector), j + border}, [2]int{-1, -1}})
		}
		if onBoard(i-(border*p.PieceVector), j-border) {
			moves = append(moves, Move{[2]int{i, j}, [2]int{i - (border * p.PieceVector), j - border}, [2]int{-1, -1}})
		}
	}

	return moves
}

func (p *Player) getValidMovesFor(pos [2]int) ([]Move, bool) {
	var result []Move
	moves, cap := p.getValidMoves()
	for _, move := range moves {
		if move.from == pos {
			result = append(result, move)
		}
	}
	if len(result) == 0 {
		cap = false
	}
	return result, cap
}

func (p *Player) hasValidMoves() bool {

	for i, row := range board {
		for j := range row {
			if !belongsToPlayer(i, j, p.Number) {
				continue
			}

			movesToCheck := p.getMovesToCheckForPiece(i, j)
			for _, move := range movesToCheck {
				valid := p.isMoveToCheck(move, movesToCheck)
				if valid {
					p.addCapture(&move)
					if isPlayersPawn(move.from[0], move.from[1], p.Number) {
						if move.calcDistance() == 1 && move.capture[0] == -1 {
							return true
						}
						if move.capture[0] != -1 {
							return true
						}
					} else {
						if move.capture[0] == -1 {
							return true
						} else {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func (p *Player) getValidMoves() ([]Move, bool) {
	var valids []Move
	var captures []Move

	for i, row := range board {
		for j := range row {
			if !belongsToPlayer(i, j, p.Number) {
				continue
			}

			movesToCheck := p.getMovesToCheckForPiece(i, j)
			for _, move := range movesToCheck {
				// fmt.Println(boardToPos(move.from), "->", boardToPos(move.to))
				valid := p.isMoveToCheck(move, movesToCheck)
				if valid {
					p.addCapture(&move)
					if isPlayersPawn(move.from[0], move.from[1], p.Number) {
						if move.calcDistance() == 1 && move.capture[0] == -1 {
							valids = append(valids, move)
						}
						if move.capture[0] != -1 {
							captures = append(captures, move)
						}
					} else {
						if move.capture[0] == -1 {
							valids = append(valids, move)
						} else {
							captures = append(captures, move)
						}
					}
				}
			}
		}
	}
	if len(captures) == 0 {
		return valids, false
	}
	return captures, true
}

func (p *Player) getValidMovesWithMultipleCapture() (bool, int, []Move) {
	moves, isCapture := p.getValidMoves()
	mCaptures := 0

	if isCapture {
		temp, mCapturesMove := p.checkMultipleCaptures(moves)
		mCaptures = temp
		if mCaptures > 0 {
			moves = []Move{mCapturesMove}
		}
	}
	return isCapture, mCaptures, moves
}

func (p *Player) isMoveToCheck(move Move, possibleMoves []Move) bool {
	if !slices.Contains(possibleMoves, move) {
		return false
	}
	if !onBoard(move.to[0], move.to[1]) {
		return false
	}
	if !isEmpty(move.to[0], move.to[1]) {
		return false
	}

	return true
}

func (p *Player) addCapture(move *Move) {
	dist := move.calcDistance()

	if dist < 1 {
		return
	}

	vectVert := move.to[1] - move.from[1]
	vectHor := move.to[0] - move.from[0]
	if vectVert > 0 {
		vectVert = 1
	} else {
		vectVert = -1
	}

	if vectHor > 0 {
		vectHor = 1
	} else {
		vectHor = -1
	}
	checkHor := move.from[0]
	checkVert := move.from[1]

	for i := 0; i < dist; i++ {
		checkHor += vectHor
		checkVert += vectVert
		if belongsToEnemy(checkHor, checkVert, p.Number) {
			if isEmpty(checkHor+vectHor, checkVert+vectVert) {
				move.capture = [2]int{checkHor, checkVert}
				return
			}
			move.capture = [2]int{-1, -1}
			return
		}
	}
}

func (p *Player) executeMove(move Move) {

	king := false
	if board[move.from[0]][move.from[1]] == p.Number+kingShift {
		king = true
	} else if ((p.PieceVector == -1 && move.to[0] == 0) || (p.PieceVector == 1 && move.to[0] == 7)) && move.capture[0] == -1 {
		king = true
	}

	board[move.from[0]][move.from[1]] = 3
	if king {
		board[move.to[0]][move.to[1]] = p.Number + kingShift
	} else {
		board[move.to[0]][move.to[1]] = p.Number
	}

	if move.capture[0] != -1 { // Capture
		board[move.capture[0]][move.capture[1]] = 3
	}
}

func (p *Player) checkMultipleCaptures(moves []Move) (int, Move) {
	max := 0
	bestMove := Move{[2]int{-1, -1}, [2]int{-1, -1}, [2]int{-1, -1}}
	for _, move := range moves {
		temp := board
		p.executeMove(move)
		newMoves, isCap := p.getValidMovesFor(move.to)
		if isCap {

			tempVal, _ := p.checkMultipleCaptures(newMoves)
			tempVal++

			if tempVal >= max {
				max = tempVal
				bestMove = move
			}
		}
		board = temp
	}
	// fmt.Println(boardToPos(bestMove.from), "->", boardToPos(bestMove.to))
	return max, bestMove
}

func (p *Player) calculateCost() int {
	cost := 0
	area3 := [2][2]int{PosToBoard("c3"), PosToBoard("f6")}
	area2 := [2][2]int{PosToBoard("b2"), PosToBoard("g7")}

	for i, row := range board {
		for j := range row {
			if belongsToPlayer(i, j, p.Number) {
				if i <= area3[0][0] && i >= area3[1][0] && j >= area3[0][1] && j <= area3[1][1] {
					// fmt.Println("3:", boardToPos([2]int{i, j}))
					cost += 3
				} else if i <= area2[0][0] && i >= area2[1][0] && j >= area2[0][1] && j <= area2[1][1] {
					// fmt.Println("2:", boardToPos([2]int{i, j}))
					cost += 2
				} else {
					// fmt.Println("1:", boardToPos([2]int{i, j}))
					cost++
				}
			}
		}
	}
	return cost
}

func (p *Player) PlayerTurn() {
	for p.playerTurn(os.Stdin) {
	}
}

func (p *Player) playerTurn(ioReader io.Reader) bool {
	PrintBoard()
	reader := bufio.NewReader(ioReader)

	fmt.Println("cost:", p.calculateCost())
	fmt.Printf("Player %d (%s) move (ex: a3 b4): \n", p.Number, p.Symbol)
	moves, isCapture := p.getValidMoves()

	mCaptures, mCapturesMove := p.checkMultipleCaptures(moves)
	if isCapture && mCaptures > 0 {
		fmt.Println("multiple captures", mCaptures)
		moves = []Move{mCapturesMove}
	}

	fmt.Println("moves")
	for _, m := range moves {
		fmt.Println(boardToPos(m.from), "->", boardToPos(m.to))
	}

	text, _ := reader.ReadString('\n')
	fmt.Println()
	text = string([]rune(strings.ToLower(text))[0:5])

	match, _ := regexp.MatchString(`^[a-h][1-8] [a-h][1-8].*`, text)
	if !match {
		fmt.Printf("Bad input, expected [a-h][1-8] [a-h][1-8], got %s\n", text)
		return true
	}

	pos := strings.Split(text, " ")
	mov1 := PosToBoard(pos[0])
	mov2 := PosToBoard(pos[1])

	move := Move{mov1, mov2, [2]int{-1, -1}}

	return p.playerTurnLogic(move)
}

func (p *Player) playerTurnLogic(move Move) (repeat bool) {

	isCapture, mCaptures, moves := p.getValidMovesWithMultipleCapture()
	// fmt.Println(boardToPos(move.from), "->", boardToPos(move.to))

	for _, m := range moves {
		if m.from == move.from && m.to == move.to {
			if isPlayersKing(m.from[0], m.from[1], p.Number) {
				p.KingMoves++
			}
			p.executeMove(m)

			//do another turn when multiple captures
			if isCapture && mCaptures > 0 {
				return true
			}
			return false
		}
	}
	fmt.Println("invalid move")
	return true
}
