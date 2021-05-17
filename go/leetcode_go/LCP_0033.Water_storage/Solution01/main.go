package main

import (
	"fmt"
)

func storeWater(bucket []int, vat []int) int {
	if sum(vat) == 0 {
		return 0
	}

	bLen := len(bucket)
	res := 10000 + bLen

	for p := 1; p < 10000; p++ {
		tmp := p
		for i := 0; i < bLen; i++ {
			buckets := (vat[i]+p-1)/p - bucket[i]
			if buckets > 0 {
				tmp += buckets
			}
		}

		if tmp < res {
			res = tmp
		}
	}

	return res
}

func sum(vat []int) int {
	res := 0
	for _, v := range vat {
		res += v
	}
	return res
}

func test01() {
	bucket, vat := []int{1, 3}, []int{6, 8}
	succResult := 4
	fmt.Println("test01 bucket:=", bucket, " vat:=", vat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", storeWater(bucket, vat))
	fmt.Println()
}

func test02() {
	bucket, vat := []int{9, 0, 1}, []int{0, 2, 2}
	succResult := 3
	fmt.Println("test02 bucket:=", bucket, " vat:=", vat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", storeWater(bucket, vat))
	fmt.Println()
}

func test03() {
	bucket, vat := []int{2, 27, 24, 75}, []int{53, 52, 28, 82}
	succResult := 13
	fmt.Println("test03 bucket:=", bucket, " vat:=", vat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", storeWater(bucket, vat))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
