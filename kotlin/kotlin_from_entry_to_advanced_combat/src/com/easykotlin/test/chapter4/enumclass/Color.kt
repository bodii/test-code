package com.easykotlin.test.chapter4.enumclass

// 声明一个带构造参数的枚举类
enum class Color(val rgb: Int) {
    RED(0xFF0000),
    GREEN(0x00FF00),
    BLUE(0x0000FF)
}

fun main(args: Array<String>) {
    val c = Color.GREEN
    println(c)

    println(c.rgb)

    println(c.ordinal)

    println(c.name)

    println((0x00FF00 to Int).first)
}