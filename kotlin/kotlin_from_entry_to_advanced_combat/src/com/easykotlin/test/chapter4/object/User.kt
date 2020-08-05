package com.easykotlin.test.chapter4.`object`

/**
 * Kotlin中没有静态属性和方法，但是可以使用关键字object声明一个
 * object单例对象
 *
 */
object User {
    val username: String = "admin"
    val password: String = "admin"

    fun hello () {
        println("Hello, object!")
    }
}

fun main(args: Array<String>) {
    println(User.username)
    println(User.password)
    User.hello()
}