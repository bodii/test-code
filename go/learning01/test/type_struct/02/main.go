package main

import (
	"fmt"
	"os"
)

// 学生管理系统

// 菜单方法
func showEnum() {
	fmt.Println(`
Welcome sms!"
	1. 查看所有学生
	2. 查看某个学生是否存在
	3. 添加学生信息
	4. 修改学生信息
	5. 删除学生信息
	6. 退出

	`)
}

// 退出方法
func quit() {
	os.Exit(1)
}

func main() {
	// 实例一个学生管理者
	sm := studentManager{students: make(map[idType]student, 100)}

	for {
		showEnum()
		// 等待用户输入
		fmt.Print("请输入要执行的编号: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			sm.shows()
		case 2:
			sm.show()
		case 3:
			sm.add()
		case 4:
			sm.edit()
		case 5:
			sm.del()
		case 6:
			quit()
		default:
			fmt.Println("输入错误，请重新选择要执行的功能")
		}
	}

}
