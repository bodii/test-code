package com.easykotlin.test.chapter9.threadtest

import kotlin.concurrent.thread

fun main(args: Array<String>) {
    val t = Thread{
        Thread.sleep(2000)
        println("C 使用lambda 表达式: ${Thread.currentThread()}")
    }

    t.isDaemon = false
    t.name = "CThread"
    t.priority = 3
    t.start()

    thread(start = true, isDaemon = false, name = "DThread", priority = 3) {
        Thread.sleep(1000)
        println("D 使用Kotlin ")
    }

}