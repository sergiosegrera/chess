package board

type Color bool
type Type int

const (
	WHITE        = true
	WHITE_KING   = "♔"
	WHITE_QUEEN  = "♕"
	WHITE_ROOK   = "♖"
	WHITE_BISHOP = "♗"
	WHITE_KNIGHT = "♘"
	WHITE_PAWN   = "♙"

	BLACK        = false
	BLACK_KING   = "♚"
	BLACK_QUEEN  = "♛"
	BLACK_ROOK   = "♜"
	BLACK_BISHOP = "♝"
	BLACK_KNIGHT = "♞"
	BLACK_PAWN   = "♟︎"

	PAWN   = 0
	KNIGHT = 1
	BISHOP = 2
	ROOK   = 3
	QUEEN  = 4
	KING   = 5
)

type Piece interface {
	ValidMove(*Board, *Position) bool
	ChangePosition(*Position)
	Color() Color
	Type() Type
}
