package main

import "fmt"

func main() {
	var a1 [3]bool // [true, false, true]
	fmt.Printf("a1:%T \n", a1)

	// array init
	fmt.Println(a1)

	// The first init method
	a1 = [3]bool{true, false, true}
	fmt.Println(a1)

	// The second initialization method
	a2 := [...]int{1, 23, 43, 16, 77}
	fmt.Println(a2)

	// The third initialization method
	a3 := [5]int{1, 2, 4: 1, 3: 2, 2: 7}
	fmt.Println(a3)

	// array traversal
	citys := [...]string{"shanghai", "beijing", "changchun"}

	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for k, v := range citys {
		fmt.Println(k, v)
	}

	// 多维数组
	var aa1 [3][2]int
	aa1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(aa1)

	for _, v := range aa1 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100

	fmt.Println(b1, b2)

	a4 := [...]int{1, 3, 5, 7, 8}
	s := 0
	for _, v := range a4 {
		s += v
	}
	fmt.Println(s)

	a5 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a5); i++ {
		for j := i + 1; j < len(a5); j++ {
			if a5[i]+a5[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

	a6 := a5[:3]
	a7 := a5[3:]
	a8 := a5[:]
	fmt.Println(a6, a7, a8)
	fmt.Println(len(a6), cap(a6))
	fmt.Println(len(a7), cap(a7))

	a9 := a6[:1]
	fmt.Println(a9)
}
