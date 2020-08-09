package com.easykotlin.test.chapter7

/*
    元素去重方法 dictinct()
 */
fun main(args: Array<String>) {
    val dupList = listOf(1, 1, 2, 2, 3, 3, 3)
    println(dupList.distinct())
}