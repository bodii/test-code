package com.easykotlin.test.chapter2

fun main(args: Array<String>) {
    var x = 10
    while (x > 0) {
        x --
        println(x)
    }

    var y = 10
    do {
        y ++
        println(y)
    } while(y < 20)
}
