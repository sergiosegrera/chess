package cli

import (
	"github.com/sergiosegrera/chess/board"
)

// e2-e4 -> (6, 4) (4, 4)
// Y: a -> 0 ... h -> 7
// X: 1 -> 7 ... 8 -> 0

func LongAlgebraicToPositions(s string) (*board.Position, *board.Position) {
	// Temporary
	// Unsure if best practice
	var from board.Position
	from.Y = int(s[0]) - 97
	from.X = 56 - int(s[1])
	var to board.Position
	to.Y = int(s[3]) - 97
	to.X = 56 - int(s[4])
	return &from, &to
}
