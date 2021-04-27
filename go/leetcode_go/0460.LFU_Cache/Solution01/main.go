package main

import (
	"container/list"
	"fmt"
)

type LFUCache struct {
	Cap     int
	Keys    map[int]*list.Element
	MinList map[int]*list.List
	Min     int
}

type node struct {
	K, V, Visits int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Cap:     capacity,
		Keys:    make(map[int]*list.Element),
		MinList: make(map[int]*list.List),
		Min:     0,
	}
}

func (this *LFUCache) Get(key int) int {
	el, ok := this.Keys[key]
	if !ok {
		return -1
	}
	currentNode := el.Value.(*node)
	this.MinList[currentNode.Visits].Remove(el)
	// 如果之前的点击率 等于 最小值，并且最小点击率的元素为0（之前只剩它自己）时，
	if currentNode.Visits == this.Min && this.MinList[currentNode.Visits].Len() == 0 {
		this.Min++
	}

	// 当前点击率加1
	currentNode.Visits++
	// 如果加1后的点击率key作为的列表中如果不存在值，则创建
	if _, ok := this.MinList[currentNode.Visits]; !ok {
		this.MinList[currentNode.Visits] = list.New()
	}

	// 将新增加点击率的的key，对应到修改后点击率节点，并放置到新小列表元素的第一位
	minList := this.MinList[currentNode.Visits]
	newNode := minList.PushFront(currentNode)
	// 维护Keys
	this.Keys[key] = newNode

	return currentNode.V
}

func (this *LFUCache) Put(key int, value int) {
	// 如果总容量小于1，则直接返回
	if 0 == this.Cap {
		return
	}

	// 如果当前key存在, 则更新
	if elementNode, ok := this.Keys[key]; ok {
		currentNode := elementNode.Value.(*node)
		currentNode.V = value
		this.Get(key)
		return
	}

	// 如果当前容量等于元素数,则剔除一个最小列表最后一个
	if this.Cap == len(this.Keys) {
		minCurrentList := this.MinList[this.Min]
		backNode := minCurrentList.Back()
		delete(this.Keys, backNode.Value.(*node).K)
		minCurrentList.Remove(backNode)
	}

	this.Min = 1
	// 如果容量没有超过，则建节点并将节点添加到MinList和Keys
	if _, ok := this.MinList[1]; !ok {
		this.MinList[1] = list.New()
	}

	currentNode := &node{K: key, V: value, Visits: 1}
	minList := this.MinList[1]
	newNode := minList.PushFront(currentNode)
	this.Keys[key] = newNode
}

/**
 * Your LFUCache object will be instantiated and called as such:
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
	fmt.Println("get 3 result:", c.Get(3))
	c.Put(4, 4)
	fmt.Println("get 1 result:", c.Get(1))
	fmt.Println("get 3 result:", c.Get(3))
	fmt.Println("get 4 result:", c.Get(4))

}

func main() {
	test01()
}
