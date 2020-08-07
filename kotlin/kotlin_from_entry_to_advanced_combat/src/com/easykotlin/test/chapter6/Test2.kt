package com.easykotlin.test.chapter6

import com.easykotlin.test.chapter6.firstChar as bb
import com.easykotlin.test.chapter6.lastChar

fun main(args: Array<String>) {
    val str = "abc"
    str.bb().also {
        println(it)
    }
    str.lastChar().also {
        println(it)
    }

    with(str.lastChar()) {
        println(this)
    }
}