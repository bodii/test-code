package main

import (
	"container/list"
	"fmt"
)

const base = 769

type dict struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (hm *MyHashMap) hash(key int) int {
	return key % base
}

func (hm *MyHashMap) Put(key, value int) {
	h := hm.hash(key)
	for e := hm.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(dict); et.key == key {
			e.Value = dict{key, value}
			return
		}
	}
	hm.data[h].PushBack(dict{key, value})
}

func (hm *MyHashMap) Get(key int) int {
	h := hm.hash(key)
	for e := hm.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(dict); et.key == key {
			return et.value
		}
	}
	return -1
}

func (hm *MyHashMap) Remove(key int) {
	h := hm.hash(key)
	for e := hm.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(dict).key == key {
			hm.data[h].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func test01() {
	myHashMap := Constructor()
	myHashMap.Put(1, 1)                                          // myHashMap 现在为 [[1,1]]
	myHashMap.Put(2, 2)                                          // myHashMap 现在为 [[1,1], [2,2]]
	fmt.Println("success result: 1, result:", myHashMap.Get(1))  // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
	fmt.Println("success result: -1, result:", myHashMap.Get(3)) // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Put(2, 1)                                          // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
	fmt.Println("success result: 1, result:", myHashMap.Get(2))  // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
	myHashMap.Remove(2)                                          // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
	fmt.Println("success result: -1, result:", myHashMap.Get(2)) // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
}

func main() {
	test01()
}
