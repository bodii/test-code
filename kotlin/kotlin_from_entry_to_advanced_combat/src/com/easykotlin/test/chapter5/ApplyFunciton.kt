package com.easykotlin.test.chapter5

fun testApply() {
    // 普通写法
    val list = mutableListOf<String>()
    list.add("A")
    list.add("B")
    list.add("C")
    println("普通写法 list = $list")
    println(list)

    // 使用apple()函数的写法
    val a = ArrayList<String>().apply {
        add("A")
        add("B")
        add("C")
        println("使用apply函数写法this=$this")
    }

    println(a)
    // 等价于
    a.let { println(it) }

    val b = mutableListOf<String>().apply {
        add("A")
        add("B")
        println(this)
    }
}

fun main(args: Array<String>) {
    testApply()
}