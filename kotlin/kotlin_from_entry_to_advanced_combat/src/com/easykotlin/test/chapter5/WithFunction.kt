package com.easykotlin.test.chapter5

fun testWithFun() {
    // 普通写法
    val list = mutableListOf<String>()
    list.add("A")
    list.add("B")
    list.add("C")
    println("常规写法 list = $list")  // 常规写法 list = [A， B， C]

    // 使用with()函数写法
    with (ArrayList<String>()) {
        add("A")
        add("B")
        add("C")
        println("使用 with 函数写法 this = $this")
    }.let {
        println(it)
    }
}

fun main(args: Array<String>) {
    testWithFun()
}