package com.easykotlin.test.chapter4

/**
 * 在Kotlin中，声明类的同时声明构造函数
 */
class Person(var name: String, var age: Int, var sex: String) {
    override fun toString(): String { // override关键字，重写toString()
        return "Person(name='$name', age=$age, sex='$sex')"
    }
}

fun main(args: Array<String>) {
    val person = Person("Jack", 29, "M")
    println("person = ${person}")
}