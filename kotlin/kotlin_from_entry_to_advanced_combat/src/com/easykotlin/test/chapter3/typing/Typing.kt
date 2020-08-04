package com.easykotlin.test.chapter3.typing

fun test001() {
    val a: Int = 1000
    val b: Int = 1000
    println("a === b :" + (a === b) )
    println("a == b :" + (a == b) )
}

/**
 * 当a、b都为可空类型时，a与b的引用是不等的。
 */
fun test002() {
    val a: Int? = 1000
    val b: Int? = 1000

    println("a == b: " + (a == b))
    println("a === b: " + (a === b))
}

fun main(args: Array<String>) {
    test001()
    test002()
}