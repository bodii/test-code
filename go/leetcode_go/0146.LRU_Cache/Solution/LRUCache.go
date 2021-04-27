package main

import "container/list"

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

func (this *LRUCache) Get(key int) int {
	if el, ok := this.Keys[key]; ok {
		this.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if el, ok := this.Keys[key]; ok {
		el.Value = pair{K: key, V: value}
		this.List.MoveToFront(el)
	} else {
		el := this.List.PushFront(pair{K: key, V: value})
		this.Keys[key] = el
	}

	if this.List.Len() > this.Cap {
		el := this.List.Back()
		this.List.Remove(el)
		delete(this.Keys, el.Value.(pair).K)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
