package com.easykotlin.test.chapter2


fun fortest01() {
    val collection = 1..10
    for(item in collection) {
        print(item)
    }
}

fun fortest02() {
    val array = listOf("Java", "Kotlin", null, "Go")
    for (i in array.indices) {
        print(array[i] + ", ")
    }
}

fun fortest03() {
    val array = 1..10
    for ((index, value) in array.withIndex()) {
        println("the element at $index is $value")
    }
}

fun fortest04() {
    for (i: Int in 1..10) {
        print("$i, ")
    }
}

fun fortest05() {
    (1..10).forEach{ print("$it, ") }
}

fun fortest06() {
    val array = 1.rangeTo(10)
    for (i: Int in array)
        print("$i, ")
}

fun sumFact(n: Int): Int {
    var sum = 0
    for (i in 1..n) {
        sum += fact(i)
    }

    return sum
}

fun main(args: Array<String>) {
    fortest01()
    println()

    fortest02()
    println()

    fortest03()
    println()

    fortest04()
    println()

    fortest05()
    println()

    fortest06()
    println()

    println("sumFact result:" + sumFact(10))
}