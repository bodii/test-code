package com.easykotlin.test.chapter2

fun main(args: Array<String>) {
    for (i in 1..10) {
        println(i)
        if (i % 2 == 0)
            break
    }

    for (i in 1..10) {
        if (i % 2 == 0)
            continue
        println(i)
    }

    println("sum a=10, b=13: ${sum(10, 13)}")
    println("max a=7, b=5: ${max2(7, 5)}")
}

fun sum(a: Int, b: Int): Int {
    return a + b
}

fun max1(a: Int, b: Int): Int {
    return if (a > b) {
        a
    } else {
        b
    }
    // return if (a > b) a else b
}