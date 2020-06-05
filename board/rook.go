package board

type Rook struct {
	pieceType Type
	code      string
	color     Color
	position  *Position
}

func NewRook(color Color, position *Position) *Rook {
	var r Rook
	r.pieceType = ROOK
	if color == WHITE {
		r.code = WHITE_ROOK
	} else {
		r.code = BLACK_ROOK
	}
	r.color = color
	r.position = position
	return &r
}

func (r *Rook) ValidMove(board *Board, position *Position) bool {
	if position.X > 7 || position.X < 0 || position.Y > 7 || position.Y < 0 {
		return false
	}

	deltaX := r.position.X - position.X
	deltaY := r.position.Y - position.Y
	if deltaX == 0 && deltaY == 0 {
		return false
	}
	// up
	if deltaX > 0 {
		if deltaY != 0 {
			return false
		}
		for i := r.position.X - 1; i > position.X; i-- {
			if !board.IsEmpty(&Position{X: i, Y: r.position.Y}) {
				return false
			}
		}
	}
	// down
	if deltaX < 0 {
		if deltaY != 0 {
			return false
		}
		for i := r.position.X + 1; i < position.X; i++ {
			if !board.IsEmpty(&Position{X: i, Y: r.position.Y}) {
				return false
			}
		}
	}
	// right
	if deltaY > 0 {
		if deltaX != 0 {
			return false
		}
		for i := r.position.Y - 1; i > position.Y; i-- {
			if !board.IsEmpty(&Position{X: r.position.X, Y: i}) {
				return false
			}
		}
	}
	// left
	if deltaY < 0 {
		if deltaX != 0 {
			return false
		}
		for i := r.position.Y + 1; i < position.Y; i++ {
			if !board.IsEmpty(&Position{X: r.position.X, Y: i}) {
				return false
			}
		}
	}

	piece := board.GetPiece(position)
	if piece != nil {
		if piece.Color() != !r.color {
			return false
		}
	}

	return true
}

func (r *Rook) ChangePosition(position *Position) {
	r.position = position
}

func (r *Rook) Color() Color {
	return r.color
}

func (r *Rook) Type() Type {
	return r.pieceType
}
