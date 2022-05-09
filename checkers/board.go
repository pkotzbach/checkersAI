package checkers

import (
	"fmt"
	"strconv"
	"strings"
)

const kingShift = 6
const boardSize = 8

var board [boardSize][boardSize]int

func isEmpty(i, j int) bool {
	return onBoard(i, j) && board[i][j] == 3
}

func belongsToPlayer(i, j, p int) bool {
	return isPlayersKing(i, j, p) || isPlayersPawn(i, j, p)
}

func belongsToEnemy(i, j, p int) bool {
	return !belongsToPlayer(i, j, p) && !isEmpty(i, j)
}

func isPlayersPawn(i, j, p int) bool {
	return board[i][j] == p
}

func isPlayersKing(i, j, p int) bool {
	return board[i][j] == p+kingShift
}

func PopulateNewBoard() {
	board = [boardSize][boardSize]int{
		{0, 2, 0, 2, 0, 2, 0, 2},
		{3, 0, 2, 0, 3, 0, 2, 0},
		{0, 2, 0, 2, 0, 2, 0, 2},
		{3, 0, 3, 0, 1, 0, 3, 0},
		{0, 3, 0, 3, 0, 1, 0, 3},
		{2, 0, 1, 0, 1, 0, 1, 0},
		{0, 3, 0, 3, 0, 3, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0}}
	//f4 g5
	// {0, 2, 0, 2, 0, 2, 0, 2},
	// {2, 0, 2, 0, 2, 0, 2, 0},
	// {0, 2, 0, 2, 0, 2, 0, 2},
	// {3, 0, 3, 0, 3, 0, 3, 0},
	// {0, 3, 0, 3, 0, 3, 0, 3},
	// {1, 0, 1, 0, 1, 0, 1, 0},
	// {0, 1, 0, 1, 0, 1, 0, 1},
	// {1, 0, 1, 0, 1, 0, 1, 0}}
}

func PrintBoard() {
	fmt.Println(BoardAsString())
}

func BoardAsString() string {
	pieces := map[int]string{
		0: "=",
		1: "X",
		2: "O",
		3: " ",
		7: "%",
		8: "0",
	}
	boardString := ""
	for i := 0; i < boardSize; i++ {
		boardString += strconv.Itoa(8 - i)
		for j := 0; j < boardSize; j++ {
			boardString += " | " + pieces[board[i][j]]
		}
		boardString += " |\n"

	}
	boardString += "    A   B   C   D   E   F   G   H "
	return boardString
}

func onBoard(vert int, hoz int) bool {
	if vert >= boardSize || vert < 0 || hoz >= boardSize || hoz < 0 {
		return false
	}
	return true
}

func PosToBoard(a string) [2]int {
	cvt := []rune(strings.ToLower(a))
	hoz := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
	}
	vert := map[string]int{
		"8": 0,
		"7": 1,
		"6": 2,
		"5": 3,
		"4": 4,
		"3": 5,
		"2": 6,
		"1": 7,
	}
	h := hoz[strings.ToLower(string(cvt[0]))]
	v := vert[strings.ToLower(string(cvt[1]))]

	return [2]int{v, h}
}

func boardToPos(pos [2]int) string {
	vert := map[int]string{
		0: "a",
		1: "b",
		2: "c",
		3: "d",
		4: "e",
		5: "f",
		6: "g",
		7: "h",
	}

	return vert[pos[1]] + strconv.Itoa(8-pos[0])
}

func EnemyOnBoard(player int) bool {
	enemy := 2
	if player == 2 {
		enemy = 1
	}

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == enemy || board[i][j] == enemy+6 {
				return true
			}
		}
	}
	return false
}
