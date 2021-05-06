package main

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	left, right := 0, 0
	len := len(distance)

	for i := start; i != destination; i = (i + 1) % len {
		left += distance[i]
	}

	for i := destination; i != start; i = (i + 1) % len {
		right += distance[i]
	}

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
	fmt.Printf("test04 distance:= %v, start:= %d, desination:=%d\n",
		distance, start, desination)
	fmt.Printf("success result: [ %d ]\n", 17)
	fmt.Println("result:", distanceBetweenBusStops(distance, start, desination))
}

func test05() {
	distance := []int{3, 6, 7, 2, 9, 10, 7, 16, 11}
	start, desination := 6, 2
	fmt.Printf("test05 distance:= %v, start:= %d, desination:=%d\n",
		distance, start, desination)
	fmt.Printf("success result: [ %d ]\n", 28)
	fmt.Println("result:", distanceBetweenBusStops(distance, start, desination))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
