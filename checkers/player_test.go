package checkers

import (
	"fmt"
	"strings"
	"testing"
)

func execMoveForTests(p *Player, move *Move) {
	p.playerTurnLogic(*move)
}

func TestPlayer_Rule3(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 2, 0, 2, 0, 2},
		{1, 0, 2, 0, 2, 0, 2, 0},
		{0, 3, 0, 2, 0, 2, 0, 2},
		{3, 0, 2, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0}}

	expectedBoard :=
		`8 | = |   | = | O | = | O | = | O |
		7 | X | = | O | = | O | = | O | = |
		6 | = |   | = | O | = | O | = | O |
		5 |   | = | O | = |   | = |   | = |
		4 | = |   | = |   | = |   | = |   |
		3 | X | = | X | = | X | = | X | = |
		2 | = | X | = | X | = | X | = | X |
		1 | X | = | X | = | X | = | X | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	// p2 := Player{Number: 2, Symbol: "O", PieceVector: 1}

	execMoveForTests(&p1, &Move{from: PosToBoard("a7"), to: PosToBoard("a8")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule3 failed!\nexpected:")
		t.Error(exp)
		t.Error("got")
		t.Error(act)
	}
}

func TestPlayer_Rule7(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 2, 0, 2, 0, 2},
		{1, 0, 2, 0, 2, 0, 2, 0},
		{0, 3, 0, 2, 0, 2, 0, 2},
		{3, 0, 2, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0}}

	expectedBoard :=
		`8 | = |   | = | O | = | O | = | O |
		7 | X | = | O | = | O | = | O | = |
		6 | = |   | = | O | = | O | = | O |
		5 |   | = | O | = |   | = |   | = |
		4 | = | X | = |   | = |   | = |   |
		3 |   | = | X | = | X | = | X | = |
		2 | = | X | = | X | = | X | = | X |
		1 | X | = | X | = | X | = | X | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	// p2 := Player{Number: 2, Symbol: "O", PieceVector: 1}

	execMoveForTests(&p1, &Move{from: PosToBoard("a7"), to: PosToBoard("b6")})
	execMoveForTests(&p1, &Move{from: PosToBoard("e1"), to: PosToBoard("f2")})
	execMoveForTests(&p1, &Move{from: PosToBoard("e3"), to: PosToBoard("g5")})
	execMoveForTests(&p1, &Move{from: PosToBoard("a3"), to: PosToBoard("b4")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule7 failed!\nexpected:")
		t.Error(exp)
		t.Error("got")
		t.Error(act)
	}
}

func TestPlayer_Rule8(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 2, 0, 2, 0, 2},
		{1, 0, 2, 0, 2, 0, 2, 0},
		{0, 3, 0, 2, 0, 2, 0, 2},
		{3, 0, 2, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 1, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 2},
		{1, 0, 1, 0, 1, 0, 3, 0}}

	expectedBoard :=
		`8 | = | % | = | O | = | O | = | O |
		7 |   | = | O | = | O | = | O | = |
		6 | = |   | = | O | = | O | = | O |
		5 |   | = | O | = |   | = |   | = |
		4 | = |   | = |   | = | X | = |   |
		3 | X | = | X | = | X | = | X | = |
		2 | = | X | = | X | = | X | = |   |
		1 | X | = | X | = | X | = | 0 | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	p2 := Player{Number: 2, Symbol: "O", PieceVector: 1}

	execMoveForTests(&p1, &Move{from: PosToBoard("a7"), to: PosToBoard("b8")})
	execMoveForTests(&p2, &Move{from: PosToBoard("h2"), to: PosToBoard("g1")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule8 failed!\nexpected:")
		t.Error(exp)
		t.Error("got")
		t.Error(act)
	}
}

func TestPlayer_Rule9(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 8, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 7, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0}}

	expectedBoard :=
		`8 | = |   | = |   | = | 0 | = |   |
		7 |   | = |   | = |   | = |   | = |
		6 | = |   | = |   | = |   | = | % |
		5 |   | = |   | = |   | = |   | = |
		4 | = |   | = |   | = |   | = |   |
		3 |   | = |   | = |   | = |   | = |
		2 | = |   | = |   | = |   | = |   |
		1 |   | = |   | = |   | = |   | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	p2 := Player{Number: 2, Symbol: "O", PieceVector: 1}

	execMoveForTests(&p1, &Move{from: PosToBoard("e3"), to: PosToBoard("h6")})
	execMoveForTests(&p2, &Move{from: PosToBoard("d6"), to: PosToBoard("f8")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule9 failed!\nexpected:")
		t.Error(exp)
		t.Error("got")
		t.Error(act)
	}
}

func TestPlayer_Rule9_capture(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 2, 0, 2, 0, 2},
		{1, 0, 2, 0, 2, 0, 2, 0},
		{0, 2, 0, 3, 0, 2, 0, 2},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 2, 0, 2, 0, 3},
		{1, 0, 1, 0, 2, 0, 7, 0},
		{0, 1, 0, 1, 0, 7, 0, 3},
		{1, 0, 1, 0, 1, 0, 3, 0}}

	expectedBoard :=
		`8 | = |   | = | O | = | O | = | O |
		7 | X | = | O | = | O | = | O | = |
		6 | = | O | = |   | = | O | = | O |
		5 |   | = |   | = | % | = |   | = |
		4 | = |   | = | O | = |   | = |   |
		3 | X | = | X | = | O | = |   | = |
		2 | = | X | = | X | = | % | = |   |
		1 | X | = | X | = | X | = |   | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}

	execMoveForTests(&p1, &Move{from: PosToBoard("g3"), to: PosToBoard("b8")})
	execMoveForTests(&p1, &Move{from: PosToBoard("g3"), to: PosToBoard("e5")})
	execMoveForTests(&p1, &Move{from: PosToBoard("f2"), to: PosToBoard("c5")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule9 failed!\nexpected:")
		t.Error("\n" + exp)
		t.Error("got")
		t.Error("\n" + act)
	}
}

func TestPlayer_Rule10And11(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 2, 0, 2, 0, 2},
		{1, 0, 2, 0, 2, 0, 2, 0},
		{0, 2, 0, 2, 0, 2, 0, 2},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 2, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 3},
		{1, 0, 1, 0, 1, 0, 3, 0}}

	expectedBoard :=
		`8 | = |   | = | O | = | O | = | O |
		7 |   | = | O | = | O | = | O | = |
		6 | = |   | = | O | = | O | = | O |
		5 |   | = | X | = |   | = |   | = |
		4 | = |   | = |   | = |   | = |   |
		3 | X | = | X | = | X | = |   | = |
		2 | = | X | = | X | = | X | = | O |
		1 | X | = | X | = | X | = |   | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	p2 := Player{Number: 2, Symbol: "O", PieceVector: 1}

	execMoveForTests(&p1, &Move{from: PosToBoard("a7"), to: PosToBoard("b8")})
	execMoveForTests(&p1, &Move{from: PosToBoard("a7"), to: PosToBoard("c5")})
	execMoveForTests(&p2, &Move{from: PosToBoard("f4"), to: PosToBoard("h2")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule 10 or 11 failed!\nexpected:")
		t.Error(exp)
		t.Error("got")
		t.Error(act)
	}
}

func TestPlayer_Rule12_multipleCaptures(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 2, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 2, 0, 3, 0, 3, 0, 3},
		{1, 0, 3, 0, 3, 0, 3, 0},
		{0, 2, 0, 2, 0, 2, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0}}

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	moves, _ := p1.getValidMoves()
	a, check := p1.checkMultipleCaptures(moves)
	fmt.Println(a)
	if check.from != PosToBoard("a3") || check.to != PosToBoard("c1") || check.capture != PosToBoard("b2") {
		t.Error("multiple captures failed!")
	}
}

func TestPlayer_Rule12_noMultipleCaptures(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 2, 0, 3, 0, 3, 0, 3},
		{1, 0, 3, 0, 3, 0, 3, 0},
		{0, 2, 0, 3, 0, 2, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0}}

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	moves, _ := p1.getValidMoves()
	a, check := p1.checkMultipleCaptures(moves)
	fmt.Println(a)
	if check.from[0] != -1 {
		t.Error("no multiple captures failed!")
	}
}

func TestPlayer_Rule12_multipleCapturesWithKing(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 2},
		{3, 0, 2, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 2, 0, 3},
		{7, 0, 3, 0, 3, 0, 3, 0},
		{0, 2, 0, 2, 0, 2, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0}}

	expectedBoard :=
		`8 | = |   | = |   | = |   | = |   |
		7 |   | = |   | = |   | = |   | = |
		6 | = |   | = |   | = |   | = | O |
		5 |   | = |   | = |   | = |   | = |
		4 | = |   | = |   | = |   | = |   |
		3 |   | = |   | = |   | = |   | = |
		2 | = |   | = |   | = |   | = |   |
		1 | % | = |   | = |   | = |   | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}

	execMoveForTests(&p1, &Move{from: PosToBoard("a3"), to: PosToBoard("d6")})
	execMoveForTests(&p1, &Move{from: PosToBoard("d6"), to: PosToBoard("g3")})
	execMoveForTests(&p1, &Move{from: PosToBoard("g3"), to: PosToBoard("e1")})
	execMoveForTests(&p1, &Move{from: PosToBoard("e1"), to: PosToBoard("c3")})
	execMoveForTests(&p1, &Move{from: PosToBoard("c3"), to: PosToBoard("a1")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule 10 or 11 failed!\nexpected:")
		t.Error(exp)
		t.Error("got")
		t.Error(act)
	}
}

func TestPlayer_Rule13(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 2, 0, 2, 0, 2},
		{1, 0, 2, 0, 2, 0, 2, 0},
		{0, 2, 0, 1, 0, 2, 0, 2},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 2, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 3},
		{1, 0, 1, 0, 1, 0, 3, 0}}

	expectedBoard :=
		`8 | = | X | = | O | = | O | = | O |
		7 | X | = |   | = | O | = | O | = |
		6 | = | O | = |   | = | O | = | O |
		5 |   | = |   | = |   | = |   | = |
		4 | = |   | = |   | = | O | = |   |
		3 | X | = | X | = | X | = | X | = |
		2 | = | X | = | X | = | X | = |   |
		1 | X | = | X | = | X | = |   | = |
			A   B   C   D   E   F   G   H`

	p1 := Player{Number: 1, Symbol: "X", PieceVector: -1}
	// p2 := Player{Number: 2, Symbol: "O", PieceVector: 1}

	execMoveForTests(&p1, &Move{from: PosToBoard("d6"), to: PosToBoard("b8")})

	act := strings.ReplaceAll(strings.ReplaceAll(BoardAsString(), " ", ""), "	", "")
	exp := strings.ReplaceAll(strings.ReplaceAll(expectedBoard, " ", ""), "	", "")
	if act != exp {
		t.Error("rule13 failed!\nexpected:")
		t.Error(exp)
		t.Error("got")
		t.Error(act)
	}
}

func TestPlayer_Rule14_noMoves(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 2, 0, 3, 0, 3, 0},
		{0, 1, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 2, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0}}

	game := Game{}
	game.InitPlayers()

	execMoveForTests(&game.players[0], &Move{from: PosToBoard("b6"), to: PosToBoard("d8")})
	execMoveForTests(&game.players[1], &Move{from: PosToBoard("e3"), to: PosToBoard("f2")})

	if game.playing() {
		t.Error("rule14 failed!\nexpected:")
	}
}

func TestPlayer_Rule14_noPaws(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 2, 0, 3, 0, 3},
		{3, 0, 1, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0}}

	game := Game{}
	game.InitPlayers()

	execMoveForTests(&game.players[0], &Move{from: PosToBoard("c5"), to: PosToBoard("e7")})

	if game.playing() {
		t.Error("rule14 failed!\nexpected:")
	}
}

func TestPlayer_Rule15(t *testing.T) {
	board = [boardSize][boardSize]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 7, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 8, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0}}

	game := Game{}
	game.InitPlayers()
	var p1, p2 bool

	for i := 0; i < 8; i++ {
		game.players[0].playerTurn(strings.NewReader("c5 b4"))
		game.players[1].playerTurn(strings.NewReader("f4 g5"))
		p1 = game.players[0].playerTurn(strings.NewReader("b4 c5"))
		p2 = game.players[1].playerTurn(strings.NewReader("g5 f4"))
	}

	if p1 || p2 {
		t.Error("rule15 failed! game is on")
		t.Error("p1 last playedTurn returned", p1)
		t.Error("p2 last playedTurn returned", p2)
	}
}
