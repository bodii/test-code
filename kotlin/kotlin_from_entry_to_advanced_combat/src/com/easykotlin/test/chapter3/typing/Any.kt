package com.easykotlin.test.chapter3.typing

/**
 * Any是非空类型层次结构的根
 * Any?是可空类型层次的根。
 * Any?是Any的超集，Any?是Kotlin类型层次结构的最顶端。
 */

fun anyTest001() {
    println(1 is Any)

    println(1 is Any?)

    println(null is Any)

    println(null is Any?)

    println(Any() is Any?) // Any是Any?类型

}


fun main(args: Array<String>) {
    anyTest001()
}