package board

import "fmt"

type Position struct {
	X int
	Y int
}

type Board struct {
	Tiles [8][8]Piece
}

func NewBoard() *Board {
	var tiles [8][8]Piece
	for l := range tiles[6] {
		tiles[6][l] = NewPawn(WHITE, &Position{X: 6, Y: l})
	}
	tiles[7][0] = NewRook(WHITE, &Position{X: 7, Y: 0})
	tiles[7][1] = NewKnight(WHITE, &Position{X: 7, Y: 1})
	tiles[7][2] = NewBishop(WHITE, &Position{X: 7, Y: 2})
	tiles[7][3] = NewQueen(WHITE, &Position{X: 7, Y: 3})
	tiles[7][4] = NewKing(WHITE, &Position{X: 7, Y: 4})
	tiles[7][5] = NewBishop(WHITE, &Position{X: 7, Y: 5})
	tiles[7][6] = NewKnight(WHITE, &Position{X: 7, Y: 6})
	tiles[7][7] = NewRook(WHITE, &Position{X: 7, Y: 7})
	for l := range tiles[1] {
		tiles[1][l] = NewPawn(BLACK, &Position{X: 1, Y: l})
	}
	tiles[0][0] = NewRook(BLACK, &Position{X: 0, Y: 0})
	tiles[0][1] = NewKnight(BLACK, &Position{X: 0, Y: 1})
	tiles[0][2] = NewBishop(BLACK, &Position{X: 0, Y: 2})
	tiles[0][3] = NewQueen(BLACK, &Position{X: 0, Y: 3})
	tiles[0][4] = NewKing(BLACK, &Position{X: 0, Y: 4})
	tiles[0][5] = NewBishop(BLACK, &Position{X: 0, Y: 5})
	tiles[0][6] = NewKnight(BLACK, &Position{X: 0, Y: 6})
	tiles[0][7] = NewRook(BLACK, &Position{X: 0, Y: 7})
	return &Board{Tiles: tiles}
}

func (b *Board) Print() {
	for i := range b.Tiles {
		fmt.Printf("%v ", 8-i)
		for l := range b.Tiles[i] {
			if b.Tiles[i][l] == nil {
				fmt.Printf("- ")
			} else {
				fmt.Printf("x ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("  a b c d e f g h")
}

func (b *Board) IsEmpty(position *Position) bool {
	if b.Tiles[position.X][position.Y] == nil {
		return true
	}
	return false
}

func (b *Board) GetPiece(position *Position) Piece {
	return b.Tiles[position.X][position.Y]
}

func (b *Board) Move(from *Position, to *Position) bool {
	if b.IsEmpty(from) {
		return false
	}

	if b.Tiles[from.X][from.Y].ValidMove(b, to) {
		b.Tiles[to.X][to.Y] = nil
		b.Tiles[to.X][to.Y] = b.Tiles[from.X][from.Y]
		b.Tiles[to.X][to.Y].ChangePosition(to)
		b.Tiles[from.X][from.Y] = nil
		return true
	}

	return false
}
