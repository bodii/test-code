package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看、新增、删除学生
*/


// 定义学生结构类型
type student struct {
	no int64
	name string
}

// 创建学生构造函数
func newStudent(no int64, name string) *student {
	return &student {
		no: no,
		name: name,
	}
}

// 声明班级学生变量体
var (
	students map[int64]*student
)


// addStudent 添加学生信息
func addStudent() {
	var (
		studentID int64
		name string
	)
	fmt.Print("请输入要添加的学生学号:")
	fmt.Scanln(&studentID)

	fmt.Print("请输入要添加的学生名字:")
	fmt.Scanln(&name)
	
	s := newStudent(studentID, name)
	students[studentID] = s
}

// showStudent 显示学生信息
func showStudents() {
	validateStudents()

	for k, v := range students {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

// delStudent 删除学生信息
func delStudent() {
	validateStudents()

	var (
		studentID int64
		s *student
	)
	fmt.Print("输入要删除的学生的学号:")
	fmt.Scan(&studentID)

	s, ok:= students[studentID]
	if !ok  {
		fmt.Println("输入的学生信息不存在!")
		return
	}

	delete(students, studentID)
	fmt.Printf("%s的信息已被删除！\n", s.name)
}

func validateStudents() {
	if 1 > len(students) {
		fmt.Println("当前学生信息为空！")
		return
	}
}

// quit 退出当前程序
func quit() {
	os.Exit(1)
}

func main() {
	// 初始化班级学生
	students = make(map[int64]*student, 10)
	
	// 循环用户操作提示
	for {
		fmt.Println(`
欢迎来到班级学生系统!
	1. 查看所有学生
	2. 新增学生信息
	3. 删除学生信息
	4. 退出
		`)

		var choice int8
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			showStudents()	
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			quit()
		default:
			fmt.Println("您输入的有误，请重新选择！")
		}
	}
}
