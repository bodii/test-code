package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	K, V int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if el, ok := l.Keys[key]; ok {
		l.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if el, ok := l.Keys[key]; ok {
		el.Value = pair{K: key, V: value}
		l.List.MoveToFront(el)
	} else {
		el := l.List.PushFront(pair{K: key, V: value})
		l.Keys[key] = el
	}

	if l.List.Len() > l.Cap {
		el := l.List.Back()
		l.List.Remove(el)
		delete(l.Keys, el.Value.(pair).K)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func test01() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println("get 1 result:", c.Get(1))
	c.Put(3, 3)
	fmt.Println("get 2 result:", c.Get(2))
	c.Put(4, 4)
	fmt.Println("get 1 result:", c.Get(1))
	fmt.Println("get 3 result:", c.Get(3))
	fmt.Println("get 4 result:", c.Get(4))

}

func main() {
	test01()
}
