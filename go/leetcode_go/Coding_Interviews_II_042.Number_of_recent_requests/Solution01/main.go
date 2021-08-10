package main

import (
	"fmt"
)

type RecentCounter struct {
	ts []int
}

func Constructor() RecentCounter {
	return RecentCounter{make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.ts = append(this.ts, t)
	diff := t - 3000
	for _, t := range this.ts {
		if t < diff {
			this.ts = this.ts[1:]
		} else {
			break
		}
	}

	return len(this.ts)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func test01() {
	recentCounter := Constructor()
	fmt.Println("result success: 1, result:=", recentCounter.Ping(1))
	fmt.Println("result success: 2, result:=", recentCounter.Ping(100))
	fmt.Println("result success: 3, result:=", recentCounter.Ping(3001))
	fmt.Println("result success: 3, result:=", recentCounter.Ping(3002))

}

func main() {
	test01()
}
