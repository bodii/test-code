package com.easykotlin.test.chapter5

fun main(args: Array<String>) {
    val list = listOf(1, 2, 3, 4, 5, 6, 7)
    // 过滤函数， 参数是一个lambda表达式
    println(list.filter { it % 2 == 1 })
}