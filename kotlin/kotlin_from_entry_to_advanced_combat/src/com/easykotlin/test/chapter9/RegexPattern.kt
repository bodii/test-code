package com.easykotlin.test.chapter9

fun main(args: Array<String>) {
    val r1 = Regex("[a-z]+")
    val r2 = Regex("[a-z]+", RegexOption.IGNORE_CASE)
    var r3 = "[A-Z]+".toRegex()

    println(r1.matches("ABCzxc"))

    println(r2.matches("ABCzxc"))

    println(r3.matches("GGMM"))
}