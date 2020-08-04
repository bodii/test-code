package com.easykotlin.test.chapter2.operation

fun main(args: Array<String>) {
    test001()
    test002()
    test003()
    test004()
}

fun test001() {
    val x = null
    val y = x ?: 0
    println(y)
}

fun test002() {
    val x = false
    val y = x ?: 0
    println(y)
}

fun test003() {
    val x = ""
    val y = x ?: 0
    println(y)
}

fun test004() {
    val x = "abc"
    val y = x ?: 0
    println(y)
}
