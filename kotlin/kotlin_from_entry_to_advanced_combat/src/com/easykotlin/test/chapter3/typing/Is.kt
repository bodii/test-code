package com.easykotlin.test.chapter3.typing

/**
 * is操作符或其否定形式!is来检查对象是否符合给定类型。
 */

fun IsTest001() {
    println("abc" is String)

    println("abc" !is String)
}

fun isAny001() {
    println(null is Any)
    println(null is Any?)
}

open class Foo
class Goo : Foo()

fun main(args: Array<String>) {
    IsTest001()

    isAny001()

    val foo = Foo()
    val goo = Goo()
    println(foo is Foo)
    println(goo is Foo)
    println(foo is Goo)
    println(goo is Goo)
}
