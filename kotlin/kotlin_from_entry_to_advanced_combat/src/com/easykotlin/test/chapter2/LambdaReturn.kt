package com.easykotlin.test.chapter2

fun main(args: Array<String>) {
    println("intArrayTest01: ")
    intArrayTest01()

    println("intArrayTest02: ")
    intArrayTest02()
}

fun intArrayTest01() {
    val intArray = intArrayOf(1, 2, 3, 4, 5) // 声明一个Int数组
    intArray.forEach {
        if (it == 3) return // 在Lambda表达式中的return直接返回最近的外层函数
        println(it)
    }
}

fun intArrayTest02() {
    val intArray = intArrayOf(1, 2, 3, 4, 5)
    intArray.forEach(fun(a: Int) { // 这是一个匿名函数
        if (a == 3) return
        // 循环会继续
        println(a)
    })
}