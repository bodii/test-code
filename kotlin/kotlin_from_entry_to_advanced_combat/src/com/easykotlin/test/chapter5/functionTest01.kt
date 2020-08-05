package com.easykotlin.test.chapter5

/**
 * Kotlin 是一种面向表达式的语言。
 */

val sum = fun(x: Int, y: Int): Int {return x + y}

fun multiply(x: Int, y: Int): Int {
    return x * y
}

fun main(args: Array<String>) {
    println(sum)

    println(sum(1, 1))

    println(multiply(1, 2))
}