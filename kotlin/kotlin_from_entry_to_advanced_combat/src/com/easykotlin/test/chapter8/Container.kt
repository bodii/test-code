package com.easykotlin.test.chapter8

class Container<K, V>(var key: K, var value: V) {
    override fun toString(): String {
        return "Container(key=$key, value=$value)"
    }
}

fun testContainer() {
    val container = Container<Int, String>(1, "A")
    println(container)
}

fun main(args: Array<String>) {
    testContainer()
}