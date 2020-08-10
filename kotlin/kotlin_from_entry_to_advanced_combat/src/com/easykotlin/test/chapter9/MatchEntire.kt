package com.easykotlin.test.chapter9

fun main(args: Array<String>) {
    val re = Regex("[0-9]+")
    println(re.matchEntire("1234567890"))

    println(re.matchEntire("1234567890!"))

    println(re.matchEntire("1234567890")?.value)
}