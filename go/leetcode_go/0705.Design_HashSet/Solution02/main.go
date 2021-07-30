package main

import "fmt"

type MyHashSet struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]int, 1000001)}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.data[key] = 0
}

/** Returns true if this set Contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.data[key] == 1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func test01() {
	myHashSet := Constructor()
	myHashSet.Add(1)                   // set = [1]
	myHashSet.Add(2)                   // set = [1, 2]
	fmt.Println(myHashSet.Contains(1)) // 返回 True
	fmt.Println(myHashSet.Contains(3)) // 返回 False ，（未找到）
	myHashSet.Add(2)                   // set = [1, 2]
	fmt.Println(myHashSet.Contains(2)) // 返回 True
	myHashSet.Remove(2)                // set = [1]
	fmt.Println(myHashSet.Contains(2)) // 返回 False ，（已移除）
}

func main() {
	test01()
}
