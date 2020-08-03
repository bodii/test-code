package com.easykotlin.test.chapter2

fun main(args: Array<String>) {
    println("sum 1,10: " + sum2(1, 10))
    println("max 1,2: " + max2(1, 2))
}

// Kotlin中可以直接使用“=”符号返回一个函数的值，这样的函数称为『函数字面量』
fun sum2(a: Int, b: Int): Int = a + b
fun max2(a: Int, b: Int): Int = if (a > b) a else b
