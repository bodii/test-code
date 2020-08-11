package com.easykotlin.test.chapter10

class Hello {
    operator fun invoke(name: String) {
        println("Hello, $name")
    }
}

fun main(args: Array<String>) {
    val hello = Hello()
    hello("World")

    hello("Kotlin")
}