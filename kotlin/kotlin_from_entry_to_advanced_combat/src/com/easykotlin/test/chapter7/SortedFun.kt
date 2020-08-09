package com.easykotlin.test.chapter7

/*
    sorted()升序函数
 */
fun main(args: Array<String>) {
    val list = listOf(1, 2, 3, 4, 5, 6, 7)
    println(list.sorted())

    val set = setOf(1, 3, 2)
    println(set.sorted())
}