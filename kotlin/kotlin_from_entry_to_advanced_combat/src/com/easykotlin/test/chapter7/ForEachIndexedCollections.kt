package com.easykotlin.test.chapter7

fun main(args: Array<String>) {
    val list = listOf(1, 2, 3, 4, 7, 5)
    list.forEachIndexed { index, value ->
        println("list index = $index, value = $value")
    }

    val set = setOf(1, 3, 2, 7, 9, 4, 5, 3)
    set.forEachIndexed { index, value ->
        println("set index = $index, value = $value")
    }

    // 直接访问entries属性获取Map中所有键/值对的Set
    val map = mapOf("x" to 1, "y" to 2, "z" to 3)
    println(map)
    println(map.entries)
    map.entries.forEach {
        println("key=${it.key}, value=${it.value}")
    }
    map.entries.forEach({
        println("key=${it.key}, value=${it.value}")
    })
}