package com.easykotlin.test.chapter2

import java.lang.IllegalArgumentException

/*
    在Kotlin中throw是表达式，它的类型是特殊类型Nothing.
    该类型没有值，与C、Java语言中的void意思一样。
 */
fun main(args: Array<String>) {
    // 使用::类引用操作符，返回这个Nothing类的类型
    println(Nothing::class) // class java.lang.Void

    failTest01("XXXX")
}

fun failTest01(msg: String): Nothing {
    throw IllegalArgumentException(msg)
}

// 如果把一个throw表达式的值赋给一个变量，需要显式声明类型为Nothing
val ex: Nothing = throw Exception("YYYYYY")