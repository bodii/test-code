package main

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	left, right := 0, 0
	len := len(distance) - 1

	for i := start; i < start+destination; i++ {
		if i >= len {
			left += distance[i-len]
		} else {
			left += distance[i]
		}
	}

	for i := start + destination; i < start+destination+destination; i++ {
		if i >= len {
			right += distance[i-len]
		} else {
			right += distance[i]
		}
	}
	fmt.Println("left:", left)
	fmt.Println("right:", right)

	if left < right {
		return left
	}
	return right
}

func test01() {
	distance := []int{1, 2, 3, 4}
	start, desination := 0, 1
	fmt.Printf("test01 distance:= %v, start:= %d, desination:=%d\n",
		distance, start, desination)
	fmt.Printf("success result: [ %d ]\n", 1)
	fmt.Println("result:", distanceBetweenBusStops(distance, start, desination))
}

func test02() {
	distance := []int{1, 2, 3, 4}
	start, desination := 0, 2
	fmt.Printf("test02 distance:= %v, start:= %d, desination:=%d\n",
		distance, start, desination)
	fmt.Printf("success result: [ %d ]\n", 3)
	fmt.Println("result:", distanceBetweenBusStops(distance, start, desination))
}

func test03() {
	distance := []int{1, 2, 3, 4}
	start, desination := 0, 3
	fmt.Printf("test03 distance:= %v, start:= %d, desination:=%d\n",
		distance, start, desination)
	fmt.Printf("success result: [ %d ]\n", 4)
	fmt.Println("result:", distanceBetweenBusStops(distance, start, desination))
}

func test04() {
	distance := []int{7, 10, 1, 12, 11, 14, 5, 0}
	start, desination := 7, 2
	fmt.Printf("test03 distance:= %v, start:= %d, desination:=%d\n",
		distance, start, desination)
	fmt.Printf("success result: [ %d ]\n", 17)
	fmt.Println("result:", distanceBetweenBusStops(distance, start, desination))
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
