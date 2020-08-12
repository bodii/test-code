package com.easykotlin.test.chapter12.reflectTest

var one = 1
fun testReflectProperty() {
    println(::one.get())
    ::one.set(2)
    println(one)
}

fun main(args: Array<String>) {
    testReflectProperty()
}