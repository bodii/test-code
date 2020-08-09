package com.easykotlin.test.chapter7

val funlist:List<(Int)->Boolean> = listOf({ it -> it % 2 == 0}, {it -> it % 2 == 1})

fun main(args: Array<String>) {
    val list = listOf(1, 2, 3, 4, 5, 6, 7)
    println(list.filter(funlist[0]))
    println(list.filter(funlist[1]))
}