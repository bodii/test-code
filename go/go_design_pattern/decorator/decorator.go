package decorator

// IDraw draw interface
type IDraw interface {
	Draw() string
}

// Square struct
type Square struct{}

// Draw draw as square function
func (s Square) Draw() string {
	return "this is a square"
}

// ColorSquare square color struct
type ColorSquare struct {
	square IDraw
	color  string
}

// NewColorSquare create an color square
func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{square: square, color: color}
}

// Draw draw color square function
func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}
