package com.easykotlin.test.chapter2

/*
    在Kotlin中任何表达式都可以用标签(label)来标记。标签的格式为标识符后跟＠符号，
    如abc@、_isOk@都是有效的标签。
    我们可以用Label标签来控制return、break或continue语句的跳转(jump)行为
 */
fun labelTest() {
    val intArray = intArrayOf(1, 2, 3, 4, 5)
    intArray.forEach here@ {
        if (it == 3) return@here

        println(it)
    }
}

fun labelTest02() {
    val intArray = intArrayOf(1, 2, 3, 4, 5)
    intArray.forEach {
        // 返回＠forEach处继续下一个循环
        // 接收该Lambda表达式的函数是forEach,所以我们可以直接使用return@forEach跳转到此处执行下一轮循环
        if (it == 3) return@forEach
        println(it)
    }
}

fun main(args: Array<String>) {
    labelTest()

    labelTest02()
}
