package board

type Color bool

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
)

type Piece interface {
	ValidMove(*Board, *Position) bool
	ChangePosition(*Position)
	Color() Color
	Print() string
}
