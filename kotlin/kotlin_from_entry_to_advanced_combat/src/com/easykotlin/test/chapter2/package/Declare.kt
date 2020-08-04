package com.easykotlin.test.chapter2.`package`

// 声明包

// 包级函数
fun what() {
    println("This is WHAT ?")
}

// 一个包下面只能有一个main()函数
fun main(args: Array<String>) {
    println("Hello, World!")
}

// 包里面的类
class Motorbike {
    fun drive() {
        println("Drive The motorbike ...")
    }
}