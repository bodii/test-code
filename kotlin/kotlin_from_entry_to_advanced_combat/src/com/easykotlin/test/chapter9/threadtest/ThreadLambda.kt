package com.easykotlin.test.chapter9.threadtest

fun main(args: Array<String>) {
    Thread {
        Thread.sleep(3000)
        println("B 使用Lambda 表达式: ${Thread.currentThread()}")
    }.start()
}