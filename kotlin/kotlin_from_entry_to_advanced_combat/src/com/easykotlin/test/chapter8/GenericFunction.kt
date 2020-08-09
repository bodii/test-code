package com.easykotlin.test.chapter8

class GenericClass {
    fun <T> console(t: T) {
        println(t)
    }
}

interface GenericInterface {
    fun <T> console(t: T)
}

fun <T: Comparable<T>> gt(x: T, y: T): Boolean {
    return x > y
}

fun main(args: Array<String>) {
    println(gt(10, 20))
}