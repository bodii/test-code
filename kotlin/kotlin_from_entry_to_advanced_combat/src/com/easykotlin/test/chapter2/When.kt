package com.easykotlin.test.chapter2

import java.lang.Integer.parseInt

fun casesWhen(obj: Any?) {
    when (obj) {
        0, 1, 2, 3, 4, 5, 6, 7, 8, 9 -> println("${obj} ==> 这是一个0-9之间的数字")
        "hello" -> println("$obj ==> 这个是字符串hello")
        is Char -> println("${obj} ==> 这是一个Char类型数据")
        else -> println("${obj} ==> else类似于Java中的case-switch中的default")
    }
}

fun switch(x: Int) {
    val s = "123"
    when (x) {
        -1, 0 -> print("x == -1 or x == 0")
        1 -> print("x == 1")
        2 -> print("x == 2")
        8 -> print("x is 8")
        parseInt(s) -> println("x is 123")
        else -> {
            print("x is neither 1 nor 2")
        }
    }
}

fun switch() {
    val x = 1
    val validNumbers = arrayOf(1, 2, 3)
    when (x) {
        in 1..10 -> print("x is in the range") // 是否在范围1..10
        in validNumbers -> print("x is valid") // 是否在数据arrayOf（1，2， 3）中
        !in 10..20 -> print("x is outside the range") // 不在范围10...20中
        else -> print("none of the above") // 默认路径
    }
}

fun fact(n: Int): Int {
    var result = 1
    when (n) {
        0, 1 -> result = 1  // 当n = 0, 1的时候，返回1
        else -> result = n * fact(n -1) // 除了0,1两种情况，递归调用自己n * fact(n-1)
    }

    return result
}

fun main(args: Array<String>) {
    casesWhen(1)
    casesWhen("hello")
    casesWhen('X')
    casesWhen(null)
    println(fact(10)) // 36288800
}