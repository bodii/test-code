package main

import "fmt"

// Widget struct
type Widget struct {
	X, Y int
}

// Label struct
type Label struct {
	Widget
	Text string
}

func test01() {
	label := Label{Widget{10, 10}, "State:"}

	label.X = 11
	label.Y = 12

	label.Paint()

	button := Button{Label{Widget{10, 10}, "State:"}}
	button.Click()
}

// Button struct
type Button struct {
	Label
}

// ListBox struct
type ListBox struct {
	Widget
	Texts []string
	Index int
}

// Painter interface
type Painter interface {
	Paint()
}

// Clicker interface
type Clicker interface {
	Click()
}

// Paint function
func (label Label) Paint() {
	fmt.Printf("%p:label.Paint(%q)\n", &label, label.Text)
}

// Paint button function
func (button Button) Paint() {
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}

// Click button function
func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}

// Paint ListBox function
func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}

// Click ListBox function
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

// NewButton new Button
func NewButton(x, y int, text string) *Button {
	return &Button{Label{Widget{50, 70}, text}}
}

func test02() {
	label := Label{Widget{10, 10}, "State:"}
	button1 := Button{Label{Widget{10, 70}, "ok"}}
	button2 := NewButton(50, 70, "Cancel")

	listBox := ListBox{Widget{10, 40}, []string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println()
	}
}

func main() {
	// test01()
	test02()
}
