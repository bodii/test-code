package com.easykotlin.test.chapter2.operation

import org.junit.Test
import org.junit.runner.RunWith
import org.junit.runners.JUnit4

// 声明 Person数据类
data class Person(val name: String, val age: Int)

// 声明Person类型的中缀操作符函数
infix fun Person.grow(years: Int): Person {
    return Person(name, age + years)
}

@RunWith(JUnit4::class)
class InfixFunctionDemoTest {
    @Test
    fun testInfixPuntion() {
        val person = Person("Jack", 20)
        println(person.grow(2)) // 直接调用函数
        println(person grow 2) // 中缀表达式调用方式
    }
}