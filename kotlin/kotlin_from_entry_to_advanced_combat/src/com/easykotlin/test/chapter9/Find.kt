package com.easykotlin.test.chapter9

fun main(args: Array<String>) {
    val re = Regex("[0-9]+")
    println(re.find("123XYZ987abcd7777"))

    println(re.find("123XYZ987abcd7777")?.value)
}