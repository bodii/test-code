package main

import "fmt"

// Employee struct
type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var list = []Employee{
	{"Jack", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

// EmployeeCountIf count employee
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for _, l := range list {
		if fn(&l) {
			count++
		}
	}

	return count
}

// EmployeeFilterIn filter employee
func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for _, l := range list {
		if fn(&l) {
			newList = append(newList, l)
		}
	}

	return newList
}

// EmployeeSumIf sum employee
func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	sum := 0
	for _, l := range list {
		sum += fn(&l)
	}

	return sum
}

func test01() {
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})

	fmt.Printf("old people: %d\n", old)
}

func test02() {
	highPay := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary > 6000
	})

	fmt.Printf("high Salary people: %d\n", highPay)
}

func test03() {
	noVacation := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Vacation == 0
	})

	fmt.Printf("People no vacation: %d\n", noVacation)
}

func test04() {
	totalPay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})

	fmt.Printf("Total Salary: %d\n", totalPay)
}

func test05() {
	youngerPay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})

	fmt.Printf("30 younger salary total: %d\n", youngerPay)
}

func main() {
	test01()

	test02()

	test03()

	test04()

	test05()
}
