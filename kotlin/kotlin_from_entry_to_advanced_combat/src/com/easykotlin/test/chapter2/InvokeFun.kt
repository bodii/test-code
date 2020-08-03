package com.easykotlin.test.chapter2

// Kotlin中可以直接使用“=”符号返回一个函数的值，这样的函数称为『函数字面量』
fun sum4(a: Int, b: Int): Int = a + b
fun max4(a: Int, b: Int): Int = if (a > b) a else b
// 直接使用表达式声明函数
val sumf = fun(a: Int, b:Int) = { a + b }

fun maxf(a: Int, b: Int) = {if (a > b) a else b}

fun main(args: Array<String>) {
    println("sum 1,10: " + sum4(1, 10))
    println("max 1,2: " + max4(1, 2))
    // 使用invoke调用函数
    println("sumf 1,2: " + sumf(1, 2).invoke())
    // 使用括号()调用函数，跟invoke()函数调用等价
    println("sumf 1,2: " + sumf(1, 2)())
    println("maxf 1,2: " + maxf(1, 2)())
}
