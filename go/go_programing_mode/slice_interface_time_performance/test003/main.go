package main

import "fmt"

// Person person struct
type Person struct {
	Name   string
	Sexual string
	Age    int
}

// PrintPerson print person
func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

// Print print person
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func execPrintPerson() {
	p := Person{
		Name:   "Jack",
		Sexual: "Male",
		Age:    44,
	}

	PrintPerson(&p)
	p.Print()
}

/*
// WithName struct
type WithName struct {
	Name string
}

// Country struct
type Country struct {
	WithName
}

// City struct
type City struct {
	WithName
}

// Printable interface
type Printable interface {
	PrintStr()
}

// PrintStr func
func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

func test02() {
	c1 := Country{WithName{"China"}}
	c2 := City{WithName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}
*/

// Country struct
type Country struct {
	Name string
}

// City struct
type City struct {
	Name string
}

// Stringable interface
type Stringable interface {
	ToString() string
}

// ToString Country toString
func (c Country) ToString() string {
	return "Country = " + c.Name
}

// ToString City tostring
func (c City) ToString() string {
	return "City = " + c.Name
}

// PrintStr function
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func test03() {
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}

// Shape ...
type Shape interface {
	Sides() int
	Area() int
}

// Square ...
type Square struct {
	len int
}

// Sides function
func (s *Square) Sides() int {
	return 4
}

func test04() {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}

/*
声明一个 _ 变量（没人用），
其会把一个 nil 的空指针，从 Square 转成 Shape，
这样，如果没有实现完相关的接口方法，编译器就会报错
*/
var _ Shape = (*Square)(nil)

func test05() {

	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}

func main() {
	// execPrintPerson()
	// test02()
	// test03()
	test04()
}
