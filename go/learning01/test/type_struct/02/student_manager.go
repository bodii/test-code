package main

import (
	"fmt"
)

type idType = int64

// 学生管理系统
type student struct {
	id   string
	name string
}

// 学生管理者结构体
type studentManager struct {
	students map[idType]student
}

// 所有学生的集合
var (
	students map[idType]student
)

// 返回一个可以获取生成学生id方法的闭包方法
func createID() func() idType {
	var id idType
	return func() idType {
		id++
		return id
	}
}

// 获取一个新的学生id方法
var getID (func() idType) = createID()

// 构造方法
func newStudent(id idType, name string) (s student) {
	s = student{
		id:   fmt.Sprintf("%06d", id),
		name: name,
	}
	return s
}

// 查看某个学生信息
func (sm studentManager) show() {
	var id idType = sm.getInputID()

	if !sm.exists(id) {
		fmt.Println("该学生编号的相关信息不存在！")
		return
	}

	fmt.Printf("学号:%s 姓名:%s\n", sm.students[id].id, sm.students[id].name)
}

// 查看所有学生信息
func (sm studentManager) shows() {
	sm.validate()

	for _, stu := range sm.students {
		fmt.Printf("学号:%s 姓名:%s\n", stu.id, stu.name)
	}
}

// 添加用户信息
func (sm studentManager) add() {
	var name string
	fmt.Print("请输入要添加的学生姓名:")
	fmt.Scanln(&name)
	if len(name) < 1 {
		fmt.Println("学生姓名输入有误！")
		return
	}

	id := getID()
	sm.students[id] = newStudent(id, name)
	fmt.Println("添加成功！")
}

// 修改某个学生信息
func (sm studentManager) edit() {
	id := sm.getInputID()
	if !sm.exists(id) {
		fmt.Println("该学生编号的相关信息不存在！")
		return
	}

	var name string
	fmt.Print("请输入学生的新姓名:")
	fmt.Scanln(&name)
	if len(name) < 1 {
		fmt.Println("学生姓名输入有误！")
		return
	}

	s, _ := sm.students[id]
	s.name = name
	sm.students[id] = s
	fmt.Println("修改成功！")
}

// 删除某个学生信息
func (sm studentManager) del() {
	var id idType = sm.getInputID()

	if !sm.exists(id) {
		fmt.Println("该学生编号的相关信息不存在！")
		return
	}

	delete(sm.students, id)
	fmt.Println("删除成功！")
}

// 检测学生信息是否为空
func (sm studentManager) isEmpty() bool {
	return len(sm.students) < 1
}

// 查看某个学生是否存在
func (sm studentManager) exists(id idType) bool {
	if _, ok := sm.students[id]; !ok {
		return false
	}

	return true
}

// 验证学生信息是否为空
func (sm studentManager) validate() {
	if sm.isEmpty() {
		fmt.Println("学生信息为空！")
		return
	}
}

// 获取输入的用户编号
func (sm studentManager) getInputID() idType {
	var id idType
	fmt.Print("请输入学生id:")
	fmt.Scan(&id)

	return id
}
