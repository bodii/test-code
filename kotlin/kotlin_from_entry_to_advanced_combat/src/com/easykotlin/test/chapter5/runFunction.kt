package com.easykotlin.test.chapter5

/**
 * run()函数
 *
 */

fun myfun(): String {
    println("执行了myfun函数")
    return "这是myfun的返回值"
}

fun testRunFun() {
    myfun()

    run( { myfun() } )
    run { myfun() }
    run { println("A") }
}

fun main(args: Array<String>) {
    testRunFun()
}