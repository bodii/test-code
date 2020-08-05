package com.easykotlin.test.chapter4.dataclass

data class LoginUser(val username: String, val password: String)

fun main(args: Array<String>) {
    val loginUser = LoginUser("admin", "admin")
    val (username, password) = loginUser // 解构声明(username, password)
    println("username = $username, password = $password")
    println("username=${loginUser.username}, password=${loginUser.password}")

    val map = mapOf(1 to "A", 2 to "B", 3 to "C")
    println(map)
}