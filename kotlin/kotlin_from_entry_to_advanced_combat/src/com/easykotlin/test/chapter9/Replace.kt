package com.easykotlin.test.chapter9

fun main(args: Array<String>) {
    val re = Regex("[0-9]+")
    println(re.replace("12345xyz", "abcd"))

    val re2 = Regex("[0-9]+")
    println(re2.replace("9XYZ8") {
        (it.value.toInt() * it.value.toInt()).toString()
    })
}