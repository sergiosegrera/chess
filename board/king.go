package board

type King struct {
	pieceType Type
	code      string
	color     Color
	position  *Position
}

func NewKing(color Color, position *Position) *King {
	var k King
	k.pieceType = KING
	if color == WHITE {
		k.code = WHITE_KING
	} else {
		k.code = BLACK_KING
	}
	k.color = color
	k.position = position
	return &k
}

func (k *King) ValidMove(board *Board, position *Position) bool {
	if position.X > 7 || position.X < 0 || position.Y > 7 || position.Y < 0 {
		return false
	}

	deltaX := k.position.X - position.X
	deltaY := k.position.Y - position.Y
	if deltaX > 1 || deltaY > 1 || deltaX < -1 || deltaY < -1 {
		return false
	}

	// TODO: King Rule: Check if new position is safe

	piece := board.GetPiece(position)
	if piece != nil {
		if piece.Color() != !k.color {
			return false
		}
	}

	return true
}

func (k *King) ChangePosition(position *Position) {
	k.position = position
}

func (k *King) Color() Color {
	return k.color
}

func (k *King) Type() Type {
	return k.pieceType
}
