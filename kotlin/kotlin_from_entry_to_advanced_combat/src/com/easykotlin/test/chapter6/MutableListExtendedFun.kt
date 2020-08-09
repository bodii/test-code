package com.easykotlin.test.chapter6

var <T> MutableList<T>.firestElement: T
    get() {
        return this[0]
    }
    set(value) {
        this[0] = value
    }

var <T> MutableList<T>.lastElement: T
    get() {
        return this[this.size - 1]
    }

    set(value) {
        this[this.size - 1] = value
    }


fun main(args: Array<String>) {
    val list = mutableListOf(1, 2, 3, 4, 5, 6, 7)
    println("list = $list")
    println("list is first element: " + list.firestElement)
    println("list is last element: " + list.lastElement)

    list.firestElement = -1
    list.lastElement = -1

    println("list = $list")
    println("list is first element: " + list.firestElement)
    println("list is last element: " + list.lastElement)
}
