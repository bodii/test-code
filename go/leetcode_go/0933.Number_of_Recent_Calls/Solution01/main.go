package main

import "fmt"

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	last := len(this.queue)

	for i := 0; i < last; {
		if this.queue[i] < t-3000 {
			this.queue = append(this.queue[0:i], this.queue[i+1:]...)
			last--
		} else {
			i++
		}
	}

	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func test01() {
	obj := Constructor()
	param_1 := obj.Ping(1)
	fmt.Println(param_1)

	param_1 = obj.Ping(100)
	fmt.Println(param_1)

	param_1 = obj.Ping(3001)
	fmt.Println(param_1)

	param_1 = obj.Ping(3002)
	fmt.Println(param_1)
}

func main() {
	test01()
}
