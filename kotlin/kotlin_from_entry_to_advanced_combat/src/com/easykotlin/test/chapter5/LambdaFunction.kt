package com.easykotlin.test.chapter5

fun main(args: Array<String>) {
    val list = listOf(1, 2, 3, 4, 5, 6, 7)
    list.filter { it % 2 != 0 }

    val isOdd = { it: Int -> it % 2 == 1}
    println(isOdd)
    println(list.filter(isOdd))
}

// public inline fun <T> Iterable<T>.filter(predicate: (T) -> Boolean): L