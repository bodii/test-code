package com.easykotlin.test.chapter3.typing

fun main(args: Array<String>) {
    val square = Array(5, {i -> i * i}) // 构造5个元素数组，元素初始值是i*i
    square.forEach(::println)
}