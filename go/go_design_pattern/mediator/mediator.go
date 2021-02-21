package mediator

import (
	"fmt"
	"reflect"
)

// Input input string type
type Input string

// String input string function
func (i Input) String() string {
	return string(i)
}

// Selection select string type
type Selection string

// Selected function
func (s Selection) Selected() string {
	return string(s)
}

// Button struct
type Button struct {
	onClick func()
}

// SetOnClick set Buttion onclick function
func (b *Button) SetOnClick(f func()) {
	b.onClick = f
}

// Dialog dialog struct
type Dialog struct {
	LoginButton         *Button
	RegButton           *Button
	Selection           *Selection
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

// HandleEvent function
func (d *Dialog) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, d.Selection):
		if d.Selection.Selected() == "登录" {
			fmt.Println("select login")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
		} else if d.Selection.Selected() == "注册" {
			fmt.Println("select register")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
		}
	}
}
