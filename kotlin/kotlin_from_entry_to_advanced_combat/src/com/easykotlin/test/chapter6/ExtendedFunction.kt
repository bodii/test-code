package com.easykotlin.test.chapter6

/**
 * Kotlin中给String扩展一个firstChar函数
 */
fun String.firstChar(): String {
    if (this.length == 0)
        return ""

    return this[0].toString()
}

fun String.lastChar(): String {
    if (this.length == 0)
        return ""

    return this[this.lastIndex].toString()
}

fun main(args: Array<String>) {
    println("abc".firstChar())
    println("abc".lastChar())
}