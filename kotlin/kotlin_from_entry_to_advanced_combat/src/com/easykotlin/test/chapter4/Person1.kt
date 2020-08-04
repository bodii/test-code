package com.easykotlin.test.chapter4

/**
 * 也可以先声明属性，等构造实例对象的时候再云初始化属性值
 */
class Person1 {
    lateinit var name: String  // lateinit 关键字表示该属性延迟初始化
    var age: Int = 0
    lateinit var sex: String

    override fun toString(): String {
        return "Person1(name='$name', age=$age, sex='$sex')"
    }
}

fun main(args: Array<String>) {
    val person1 = Person1()
    person1.name = "Jack"
    person1.age = 29
    person1.sex = "M"

    println("pserson1 = ${person1}")
}