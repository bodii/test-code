package com.easykotlin.test.chapter8

import java.util.*


interface Generator<T> {
    // 接口函数中直接使用类型T
    operator fun next(): T
}

fun testGenerator() {
    val gen = object: Generator<Int> {
        override fun next(): Int {
            return Random().nextInt(10)
        }
    }

    println(gen.next())
}

fun main(args: Array<String>) {
    testGenerator()
}