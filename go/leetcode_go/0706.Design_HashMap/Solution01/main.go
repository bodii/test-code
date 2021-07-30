package main

import "fmt"

type MyHashMap struct {
	data [1000001]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{[1000001]int{}}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value + 1
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.data[key] - 1

}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.data[key] = 0
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
