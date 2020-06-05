package board

type Pawn struct {
	pieceType Type
	code      string
	color     Color
	position  *Position
}

func NewPawn(color Color, position *Position) *Pawn {
	var p Pawn
	p.pieceType = PAWN
	if color == WHITE {
		p.code = WHITE_PAWN
	} else {
		p.code = BLACK_PAWN
	}
	p.color = color
	p.position = position
	return &p
}

func (p *Pawn) ValidMove(board *Board, position *Position) bool {
	// TODO: En passent, promotion
	// If new position out of board bounds return false
	if position.X > 7 || position.X < 0 || position.Y > 7 || position.Y < 0 {
		return false
	}
	var deltaX int
	var deltaY int
	if p.color == WHITE {
		deltaX = p.position.X - position.X
		deltaY = p.position.Y - position.Y
	} else {
		deltaX = position.X - p.position.X
		deltaY = position.Y - p.position.Y
	}

	if deltaX == 1 {
		if deltaY == 1 || deltaY == -1 {
			if piece := board.GetPiece(position); piece == nil {
				return false
			} else if piece.Color() == !p.color {
				return true
			}
		} else if deltaY == 0 {
			if board.IsEmpty(position) {
				return true
			}
		}
	} else if deltaX == 2 && deltaY == 0 && p.position.X == 6 || p.position.X == 1 {
		if board.IsEmpty(position) && board.IsEmpty(&Position{X: position.X - 1, Y: position.Y}) {
			return true
		}
	}

	return false
}

func (p *Pawn) ChangePosition(position *Position) {
	p.position = position
}

func (p *Pawn) Color() Color {
	return p.color
}

func (p *Pawn) Type() Type {
	return p.pieceType
}
