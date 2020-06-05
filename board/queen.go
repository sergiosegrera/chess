package board

type Queen struct {
	pieceType Type
	code      string
	color     Color
	position  *Position
}

func NewQueen(color Color, position *Position) *Queen {
	var q Queen
	q.pieceType = QUEEN
	if color == WHITE {
		q.code = WHITE_QUEEN
	} else {
		q.code = BLACK_QUEEN
	}
	q.color = color
	q.position = position
	return &q
}

func (q *Queen) ValidMove(board *Board, position *Position) bool {
	if position.X > 7 || position.X < 0 || position.Y > 7 || position.Y < 0 {
		return false
	}

	dx := q.position.X - position.X
	dy := q.position.Y - position.Y
	if dx == 0 && dy == 0 {
		return false
	}

	adx := dx
	ady := dy
	if adx < 0 {
		adx = -adx
	}
	if ady < 0 {
		ady = -ady
	}

	if adx == ady {
		if dx > 0 && dy > 0 {
			for i := 1; i < dx; i++ {
				if !board.IsEmpty(&Position{X: q.position.X - i, Y: q.position.Y - i}) {
					return false
				}
			}
		}
		if dx > 0 && dy < 0 {
			for i := 1; i < dx; i++ {
				if !board.IsEmpty(&Position{X: q.position.X - i, Y: q.position.Y + i}) {
					return false
				}
			}
		}
		if dx < 0 && dy > 0 {
			for i := dx - 1; i > dx; i-- {
				if !board.IsEmpty(&Position{X: q.position.X - i, Y: q.position.Y + i}) {
					return false
				}
			}
		}
		if dx < 0 && dy < 0 {
			for i := dx - 1; i > dx; i-- {
				if !board.IsEmpty(&Position{X: q.position.X - i, Y: q.position.Y - i}) {
					return false
				}
			}
		}
	} else if adx > 0 && ady == 0 || ady > 0 && adx == 0 {
		// up
		if dx > 0 {
			for i := q.position.X - 1; i > position.X; i-- {
				if !board.IsEmpty(&Position{X: i, Y: q.position.Y}) {
					return false
				}
			}
		}
		if dx < 0 {
			for i := q.position.X + 1; i < position.X; i++ {
				if !board.IsEmpty(&Position{X: i, Y: q.position.Y}) {
					return false
				}
			}
		}
		if dy > 0 {
			for i := q.position.Y - 1; i > position.Y; i-- {
				if !board.IsEmpty(&Position{X: q.position.X, Y: i}) {
					return false
				}
			}
		}
		if dy < 0 {
			for i := q.position.Y + 1; i < position.Y; i++ {
				if !board.IsEmpty(&Position{X: q.position.X, Y: i}) {
					return false
				}
			}
		}
	} else {
		return false
	}

	piece := board.GetPiece(position)
	if piece != nil {
		if piece.Color() != !q.color {
			return false
		}
	}

	return true
}

func (q *Queen) ChangePosition(position *Position) {
	q.position = position
}

func (q *Queen) Color() Color {
	return q.color
}

func (q *Queen) Type() Type {
	return q.pieceType
}
