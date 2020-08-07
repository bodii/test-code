package com.easykotlin.test.chapter5

fun testAlsoFun() {
    val a = "ABC".also {
        println(it) // 输出： ABC
    }
    println(a)  // 输出:ABC

    a.let {
        println(it)  // 输出： ABC
    }
}

fun main(args: Array<String>) {
    testAlsoFun()
}