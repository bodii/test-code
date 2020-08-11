package com.easykotlin.test.chapter10

/**
 * SQL的DSL示例
 */

// Student是一个简单的数据类
data class Student(var name: String, var sex: String, var score: Int)

// 创建List的扩展函数select()
fun <E> List<E>.select(): List<E> = this

// 然后实现where() 高阶函数
fun <E> List<E>.where(predicate: (E) -> Boolean): List<E> {
    val list = this
    val result = arrayListOf<E>()
    for (e in list) {
        if (predicate(e))
            result.add(e)
    }

    return result
}

// 实现and()高阶函数
fun <E> List<E>.add(predicate: (E) -> Boolean): List<E> {
    return where(predicate)
}

fun main(args: Array<String>) {

    // students 变量的初始化值
    val students = listOf(
            Student("jack", "M", 90),
            Student("alice", "F", 70),
            Student("bob", "M", 60),
            Student("bill", "M", 80),
            Student("helen", "F", 100)
    )

    val queryResult = students.select()
            .where { it.score >= 80 }
            .add{ it.sex == "M" }

    println(queryResult)
}