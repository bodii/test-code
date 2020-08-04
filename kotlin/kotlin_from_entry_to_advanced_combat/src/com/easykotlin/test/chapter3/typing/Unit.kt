package com.easykotlin.test.chapter3.typing

/**
 * Kotlin.Unit父类型是Any。如果是一个可空的Unit?，
 * 那么父类型是Any?.
 */

fun unitExample() {
    println("Hello, Unit")
}

fun main(args: Array<String>) {
    val helloUnit = unitExample()
    println(helloUnit)

    println("helloUnit is Unit: " + (helloUnit is Unit))
}

// 空函数体，返回类型是Unit
fun unitReturn1() {

}

// 显式return
fun unitReturn2() {
    return Unit
}

// 显示声明返回类型Unit
fun unitReturn3(): Unit {

}