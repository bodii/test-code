package com.easykotlin.test.chapter9

fun main(args: Array<String>) {
    val re = Regex("[0-9]+")
    println(re.findAll("123XYZ987abcd7777"))

    re.findAll("123XYZ987abcd7777").forEach {
        println(it.value)
    }
}