package board

type Queen struct {
	code     string
	color    Color
	position *Position
}

func NewQueen(color Color, position *Position) *Queen {
	var q Queen
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

	// Fix here

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

func (q *Queen) Print() string {
	return q.code
}
