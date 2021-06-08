package main

import (
	"fmt"
)

// Employee Definition for Employee.
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	employeeMap := make(map[int]*Employee)
	sum := 0
	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}

	dfs(employeeMap, id, &sum)
	return sum
}

func dfs(employeeMap map[int]*Employee, id int, sum *int) {
	employee := employeeMap[id]
	*sum += employee.Importance
	for _, subId := range employee.Subordinates {
		dfs(employeeMap, subId, sum)
	}
}

func test01() {
	employees, id := []*Employee{{1, 5, []int{2, 3}}, {2, 3, []int{}}, {3, 3, []int{}}}, 1
	success := 11

	fmt.Println("test01  emails:=", employees, " id:=", id)
	fmt.Println("success result:", success)
	fmt.Println("result:", getImportance(employees, id))
	fmt.Println()
}

func test02() {
	employees, id := []*Employee{{1, 2, []int{5}}, {5, -3, []int{}}, {3, 3, []int{}}}, 5
	success := -3

	fmt.Println("test02  emails:=", employees, " id:=", id)
	fmt.Println("success result:", success)
	fmt.Println("result:", getImportance(employees, id))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
