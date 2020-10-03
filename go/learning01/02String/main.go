package main

import "fmt"
import "strings"


func main() {
	s1 := "1234"
	fmt.Println(s1)

	s2 := `
	你好
		中华
	人民
		共和国
	`
	fmt.Println(s2)

	fmt.Println(len(s1))

	name := "理想"
	world := "世界"
	fmt.Println( name + world )
	ss1 := fmt.Sprintf("%s %s", name, world)
	fmt.Printf(ss1)
	fmt.Printf("%s %s", name, world)

	ret := strings.Split(ss1, ",")
	fmt.Println(ret)

	fmt.Println(strings.Contains(ss1, "理想"))
	fmt.Println(strings.HasPrefix(ss1, "理想"))
	fmt.Println(strings.HasSuffix(ss1, "理想"))


	ss2 := "abcdeb"
	fmt.Println(strings.Index(ss2, "c"))
	fmt.Println(strings.LastIndex(ss2, "b"))
}
