package com.easykotlin.test.chapter4

abstract class Shape  // 声明抽象类Shape

class Rectangle: Shape() // 继承类的语法是使用冒号“:”,父类需要在这里使用构造函数进行初始化

class Circle: Shape() // 

class Triangle: Shape()



fun main(args: Array<String>) {
    // val s = Shape() // 编译不通过！不能实例化抽象类
    val r = Rectangle()
    println(r is Shape)
}