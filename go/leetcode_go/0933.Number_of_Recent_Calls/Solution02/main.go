package main

import (
	"fmt"
	"sync"
)

type RecentCounter struct {
	q *queue
}

func Constructor() RecentCounter {
	return RecentCounter{
		q: newQueue(),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.q.add(t)
	for this.q.peek() < t-3000 {
		this.q.poll()
	}

	return this.q.size()
}

type queue struct {
	data []int
	mu   sync.Mutex
}

func newQueue() *queue {
	return &queue{data: make([]int, 0)}
}

func (q *queue) add(i int) {
	q.mu.Lock()
	q.data = append(q.data, i)
	q.sort()
	q.mu.Unlock()
}

func (q *queue) sort() {
	for i := q.size() - 1; i > 0; i-- {
		if q.data[i] < q.data[i-1] {
			q.swap(i, i-1)
		} else {
			break
		}
	}
}

func (q *queue) swap(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *queue) size() int {
	return len(q.data)
}

func (q *queue) peek() int {
	q.mu.Lock()
	last := q.size() - 1
	if last < 0 {
		q.mu.Unlock()
		return -1
	}

	v := q.data[0]
	q.mu.Unlock()

	return v
}

func (q *queue) poll() {
	q.mu.Lock()
	last := q.size()
	if last <= 0 {
		q.mu.Unlock()
		return
	}

	q.data = q.data[1:last]
	q.mu.Unlock()
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
