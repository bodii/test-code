package main

import "fmt"

type OrderedStream struct {
	data []string
	ptr  int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		data: make([]string, n+1),
		ptr:  1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.data[idKey] = value

	if this.ptr < idKey {
		return []string{}
	}

	for ; this.ptr < len(this.data); this.ptr++ {
		if this.data[this.ptr] == "" {
			break
		}
	}

	return this.data[idKey:this.ptr]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

func test01() {
	os := Constructor(5)
	fmt.Printf("success result: %s,  result: %s\n", []string{}, os.Insert(3, "ccccc"))
	fmt.Printf("success result: %s,  result: %s\n", []string{"aaaaa"}, os.Insert(1, "aaaaa"))
	fmt.Printf("success result: %s,  result: %s\n", []string{"bbbbb", "ccccc"}, os.Insert(2, "bbbbb"))
	fmt.Printf("success result: %s,  result: %s\n", []string{}, os.Insert(5, "eeeee"))
	fmt.Printf("success result: %s,  result: %s\n", []string{"ddddd", "eeeee"}, os.Insert(4, "ddddd"))
	fmt.Println()
}

func main() {
	test01()
}
