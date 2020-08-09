package com.easykotlin.test.chapter6

fun <T> List<T>.filters(predicate: (T) -> Boolean): MutableList<T> {
    val result = ArrayList<T>()
    this.forEach {
        if (predicate(it)) {
            result.add(it)
        }
    }

    return result
}

fun main(args: Array<String>) {
    val list = mutableListOf(1, 2, 3, 4, 5, 6, 7)
    println(list.filters { it -> it % 2 == 1 })
}