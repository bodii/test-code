package com.easykotlin.test.chapter12.reflectTest

open class BaseContainer<T>

class Container<T: Comparable<T>>:  BaseContainer<Int> {
    val elements: MutableList<T>

    constructor(elements: MutableList<T>) : super() {
        this.elements = elements
    }

    fun sort(): Container<T> {
        elements.sort();
        return this
    }

    override fun toString(): String {
        return "Container(elementys=$elements)"
    }

}

fun main(args: Array<String>) {
    val container = Container(mutableListOf<Int>(1, 3, 2, 5, 4, 6, 7))
    val kClass = container::class
    println(kClass)

    // 要获得Java类的引用，可以直接使用javaClass这个扩展属性
    val jClass = container.javaClass // 获取 Java class对象

    // 或者或使用KClass实例的.java属性
    val jKClass = kClass.java

}