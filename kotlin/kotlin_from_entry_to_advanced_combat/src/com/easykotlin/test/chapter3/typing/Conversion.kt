package com.easykotlin.test.chapter3.typing

/**
 * 在Kotlin中，父类禁止转换为子类型
 */

fun testIs() {
    val len = strlen("abc")
    println(3)

    val lens = strlen(1)
    println(lens)
}

fun strlen(ani: Any): Int {
    if (ani is String)
        return ani.length
    else if (ani is Number)
        return ani.toString().length
    else if (ani is Char)
        return 1
    else if (ani is Boolean)
        return 1

    print("Not A String")
    return -1
}

fun main(args: Array<String>) {
    testIs()

    // as 运算符
    val foo = Foo()
    val goo = Goo()
    println(foo as? Goo)
    println(goo as Foo)
}

