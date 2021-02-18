package flyweight

// ChessPieceUnit struct
type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "車",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
}

// NewChessPieceUnit new chess piece unit
func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

// ChessPiece chess piece
type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

// ChessBoard chess board
type ChessBoard struct {
	ChessPiece map[int]*ChessPiece
}

// NewChessBoard create an new chess board
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{ChessPiece: map[int]*ChessPiece{}}
	for id := range units {
		board.ChessPiece[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

// Move move chess board function
func (c *ChessBoard) Move(id, x, y int) {
	c.ChessPiece[id].X = x
	c.ChessPiece[id].Y = y
}
