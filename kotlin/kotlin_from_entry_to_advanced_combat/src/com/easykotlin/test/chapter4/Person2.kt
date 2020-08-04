package com.easykotlin.test.chapter4

/**
 * 声明一个具有多种构造方式的类，可以使用constructor
 */
class Person2() {
    lateinit var name: String
    var age: Int = 0
    lateinit var sex: String

    constructor(name: String): this() {
        this.name = name
    }

    constructor(name: String, age: Int): this() {
        this.name = name
        this.age = age
    }

    constructor(name: String, age: Int, sex: String): this() {
        this.name = name
        this.age = age
        this.sex = sex
    }

    override fun toString(): String {
        return "Person2(name='$name', age=$age, sex='$sex')"
    }
}


fun main(args: Array<String>) {
    val person2 = Person2("Jack", 20, "M");
    println("person2 = $person2")
}