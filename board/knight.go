package board

type Knight struct {
	pieceType Type
	code      string
	color     Color
	position  *Position
}

func NewKnight(color Color, position *Position) *Knight {
	var k Knight

	k.pieceType = KNIGHT

	if color == WHITE {
		k.code = WHITE_KNIGHT
	} else {
		k.code = BLACK_KNIGHT
	}
	k.color = color
	k.position = position
	return &k
}

func (k *Knight) ValidMove(board *Board, position *Position) bool {
	if position.X > 7 || position.X < 0 || position.Y > 7 || position.Y < 0 {
		return false
	}

	deltaX := k.position.X - position.X
	deltaY := k.position.Y - position.Y
	if deltaX < 0 {
		deltaX = -deltaX
	}
	if deltaY < 0 {
		deltaY = -deltaY
	}

	if deltaX == 2 && deltaY != 1 {
		return false
	}
	if deltaY != 2 && deltaX == 1 {
		return false
	}

	piece := board.GetPiece(position)
	if piece != nil {
		if piece.Color() != !k.color {
			return false
		}
	}

	return true
}

func (k *Knight) ChangePosition(position *Position) {
	k.position = position
}

func (k *Knight) Color() Color {
	return k.color
}

func (k *Knight) Type() Type {
	return k.pieceType
}
