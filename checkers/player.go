package checkers

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
)

type Player struct {
	Number      int
	PieceVector int
	Symbol      string
	KingMoves   int
}

type Move struct {
	from     [2]int
	to       [2]int
	captures [][2]int
}

func (m *Move) calcDistance() int {
	return int(math.Abs(float64(m.to[0] - m.from[0])))
}
func containsMoveWithoutCaptures(s []Move, b Move) bool {
	for _, a := range s {
		if a.to == b.to && a.from == b.from {
			return true
		}
	}
	return false
}

func (p *Player) getMovesToCheckForPiece(i, j int) []Move {
	var moves []Move

	if isPlayersKing(i, j, p.Number) {
		canKingMove := func(i, j int, sentinel *int) bool {
			if onBoard(i, j) {
				if *sentinel > 1 {
					return false
				}
				if belongsToEnemy(i, j, p.Number) || *sentinel > 0 {
					*sentinel++
				}
				return true
			}
			return false
		}

		sentinel1 := 0
		sentinel2 := 0
		sentinel3 := 0
		sentinel4 := 0
		border := boardSize
		for k := 1; k <= border; k++ {
			if canKingMove(i+k, j+k, &sentinel1) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + k, j + k}, make([][2]int, 0)})
			}
			if canKingMove(i+k, j-k, &sentinel2) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + k, j - k}, make([][2]int, 0)})
			}
			if canKingMove(i-k, j+k, &sentinel3) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i - k, j + k}, make([][2]int, 0)})
			}
			if canKingMove(i-k, j-k, &sentinel4) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i - k, j - k}, make([][2]int, 0)})
			}
		}
	} else if isPlayersPawn(i, j, p.Number) {
		border := 2
		for k := 1; k <= border; k++ {
			if onBoard(i+(k*p.PieceVector), j+k) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + (k * p.PieceVector), j + k}, make([][2]int, 0)})
			}
			if onBoard(i+(k*p.PieceVector), j-k) {
				moves = append(moves, Move{[2]int{i, j}, [2]int{i + (k * p.PieceVector), j - k}, make([][2]int, 0)})
			}
		}
		if onBoard(i-(border*p.PieceVector), j+border) {
			moves = append(moves, Move{[2]int{i, j}, [2]int{i - (border * p.PieceVector), j + border}, make([][2]int, 0)})
		}
		if onBoard(i-(border*p.PieceVector), j-border) {
			moves = append(moves, Move{[2]int{i, j}, [2]int{i - (border * p.PieceVector), j - border}, make([][2]int, 0)})
		}
	}

	return moves
}

func (p *Player) getPrevalidMovesFor(pos [2]int) ([]Move, bool) {
	var result []Move
	moves, _ := p.getPrevalidMoves()
	for _, move := range moves {
		if move.from == pos {
			result = append(result, move)
		}
	}
	cap := false
	for _, move := range result {
		if len(move.captures) > 0 {
			cap = true
		}
	}
	return result, cap
}

func (p *Player) getPrevalidMoves() ([]Move, bool) {
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
						if move.calcDistance() == 1 && len(move.captures) == 0 {
							valids = append(valids, move)
						}
						if len(move.captures) > 0 {
							captures = append(captures, move)
						}
					} else {
						if len(move.captures) == 0 {
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

func (p *Player) hasValidMoves() bool {
	moves, _ := p.getPrevalidMoves()
	return len(moves) > 0
}

func (p *Player) getValidMovesWithMultipleCapture() []Move {
	var (
		moves     []Move
		isCapture bool
	)

	moves, isCapture = p.getPrevalidMoves()

	if isCapture {
		temp := len(moves)
		p.checkMultipleCaptures(&moves, 0)
		if len(moves) > temp {
			var temp []Move
			maxM := moves[0]
			for _, move := range moves {
				if len(move.captures) > len(maxM.captures) {
					maxM = move
				}
			}
			temp = append(temp, maxM)
			for _, move := range moves {
				if len(move.captures) == len(maxM.captures) && move.to != maxM.to {
					temp = append(temp, move)
				}
			}
			moves = temp
		}
	}
	return moves
}

func (p *Player) isMoveToCheck(move Move, possibleMoves []Move) bool {
	if !containsMoveWithoutCaptures(possibleMoves, move) {
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
				move.captures = append(move.captures, [2]int{checkHor, checkVert})
			}
			break
		}
	}
}

func (p *Player) executeMove(move Move) {

	king := false
	if board[move.from[0]][move.from[1]] == p.Number+kingShift {
		king = true
	} else if ((p.PieceVector == -1 && move.to[0] == 0) || (p.PieceVector == 1 && move.to[0] == 7)) && len(move.captures) == 0 {
		king = true
	}

	board[move.from[0]][move.from[1]] = 3
	if king {
		board[move.to[0]][move.to[1]] = p.Number + kingShift
	} else {
		board[move.to[0]][move.to[1]] = p.Number
	}

	for _, cap := range move.captures {
		board[cap[0]][cap[1]] = 3
	}
}

func (p *Player) checkMultipleCaptures(currentMoves *[]Move, depth int) {
	count := len(*currentMoves)
	for i := depth; i < count; i++ {
		temp := board
		p.executeMove((*currentMoves)[i])
		moves, isCap := p.getPrevalidMovesFor((*currentMoves)[i].to)
		if isCap {
			for _, move := range moves {
				move.from = (*currentMoves)[i].from
				move.captures = append(move.captures, (*currentMoves)[i].captures...)
				*currentMoves = append(*currentMoves, move)
			}
			p.checkMultipleCaptures(currentMoves, count+1)
		}
		board = temp
	}
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

	moves := p.getValidMovesWithMultipleCapture()
	if len(moves) == 0 {
		return false
	}

	fmt.Println("cost:", p.calculateCost())
	fmt.Printf("Player %d (%s) move (ex: a3 b4): \n", p.Number, p.Symbol)

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

	move := Move{mov1, mov2, make([][2]int, 0)}

	return p.playerTurnLogic(move, moves)
}

func (p *Player) playerTurnLogic(move Move, moves []Move) bool {

	// fmt.Println(boardToPos(move.from), "->", boardToPos(move.to))

	for _, m := range moves {
		if m.from == move.from && m.to == move.to {
			if isPlayersKing(m.from[0], m.from[1], p.Number) {
				p.KingMoves++
			}
			p.executeMove(m)
			return false
		}
	}
	fmt.Println("invalid move")
	return true
}
