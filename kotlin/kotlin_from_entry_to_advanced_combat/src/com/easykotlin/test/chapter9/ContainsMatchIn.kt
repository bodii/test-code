package com.easykotlin.test.chapter9

fun main(args: Array<String>) {
    val re = Regex("[0-9]+")
    println(re.containsMatchIn("012Abc")) // true

    println(re.containsMatchIn("abc"))    // false
}