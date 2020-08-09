package com.easykotlin.test.chapter7

fun main(args: Array<String>) {
    val list = listOf(1, 2, 3, 4, 5, 6)
    val set = setOf(1, 2, 3, 4, 5, 6)
    val map = mapOf(1 to "a", 2 to "b", 3 to "c")

    println(list.map { it * it })
    println(set.map { it + 1 })
    println(map.map { it.value + "$" })

    val strlist = listOf("a", "b", "c")
    val newStrList = strlist.map { it -> listOf(it + 1, it + 2, it + 3, it + 4) }
    println(newStrList)

    // flatten()函数,效果是把嵌套的List结构“平铺”，变成一层的结构
    println(
            strlist.flatMap {
                it -> listOf(it + 1, it + 2, it + 3, it + 4)
            }
    )
}