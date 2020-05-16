package board

type Bishop struct {
	code     string
	color    Color
	position *Position
}

func NewBishop(color Color, position *Position) *Bishop {
	var b Bishop
	if color == WHITE {
		b.code = WHITE_BISHOP
	} else {
		b.code = BLACK_BISHOP
	}
	b.color = color
	b.position = position
	return &b
}

func (b *Bishop) ValidMove(board *Board, position *Position) bool {
	if position.X > 7 || position.X < 0 || position.Y > 7 || position.Y < 0 {
		return false
	}

	dx := b.position.X - position.X
	dy := b.position.Y - position.Y
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

	if adx != ady {
		return false
	}
	if dx > 0 && dy > 0 {
		for i := 1; i < dx; i++ {
			if !board.IsEmpty(&Position{X: b.position.X - i, Y: b.position.Y - i}) {
				return false
			}
		}
	}
	if dx > 0 && dy < 0 {
		for i := 1; i < dx; i++ {
			if !board.IsEmpty(&Position{X: b.position.X - i, Y: b.position.Y + i}) {
				return false
			}
		}
	}
	if dx < 0 && dy > 0 {
		for i := dx + 1; i > dx; i-- {
			if !board.IsEmpty(&Position{X: b.position.X - i, Y: b.position.Y + i}) {
				return false
			}
		}
	}
	if dx < 0 && dy < 0 {
		for i := dx + 1; i > dx; i-- {
			if !board.IsEmpty(&Position{X: b.position.X - i, Y: b.position.Y - i}) {
				return false
			}
		}
	}

	piece := board.GetPiece(position)
	if piece != nil {
		if piece.Color() != !b.color {
			return false
		}
	}

	return true
}

func (b *Bishop) ChangePosition(position *Position) {
	b.position = position
}

func (b *Bishop) Color() Color {
	return b.color
}

func (b *Bishop) Print() string {
	return b.code
}
