package com.easykotlin.test.chapter6

fun <T> List<T>.filter2(predicate: (T) -> Boolean): MutableList<T> {
    var result = ArrayList<T>();
    this.forEach {
        if (predicate(it))
            result.add(it)
    }

    return result
}