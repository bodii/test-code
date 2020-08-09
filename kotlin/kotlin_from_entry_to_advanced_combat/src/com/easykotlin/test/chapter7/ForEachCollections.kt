package com.easykotlin.test.chapter7

/*
    List\Set类继承了Iterable接口，里面扩展了forEach函数来迭代遍历元素;
    同样，Map接口中也扩展了forEach函数来迭代遍历元素
 */

fun main(args: Array<String>) {
    val list = listOf(1, 2, 3, 4, 5, 6)
    list.forEach {
        println(it)
    }

    val set = setOf(1, 2, 3, 7, 4, 6, 2, 5)
    set.forEach {
        println(it)
    }

    val map = mapOf(1 to "a", 2 to "b", 3 to "c")
    map.forEach {
        println("k = ${it.key}, v = ${it.value}")
    }



}