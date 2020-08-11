package com.easykotlin.test.chapter9.threadtest

fun main(args: Array<String>) {
    // 对象创建一个匿名类并覆盖run()方法
    object : Thread() {
        override fun run() {
            Thread.sleep(3000)
            println("A 使用Thread 对象表达式: ${Thread.currentThread()}")
        }
    }.start()
}