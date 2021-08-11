package main

import (
	"fmt"
)

type MovingAverage struct {
	data []int
	size int
	sum  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		data: make([]int, 0),
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	currLen := len(this.data)

	if currLen >= this.size {
		this.sum -= this.data[0]
		this.data = this.data[1:]
	} else {
		currLen++
	}

	this.sum += val
	this.data = append(this.data, val)

	return float64(this.sum) / float64(currLen)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func test01() {
	movingAverage := Constructor(3)
	fmt.Println("result success: 1.0, result:=", movingAverage.Next(1))
	fmt.Println("result success: 5.5, result:=", movingAverage.Next(10))
	fmt.Println("result success: 4.66667, result:=", movingAverage.Next(3))
	fmt.Println("result success: 6.0, result:=", movingAverage.Next(5))

}

func main() {
	test01()
}
