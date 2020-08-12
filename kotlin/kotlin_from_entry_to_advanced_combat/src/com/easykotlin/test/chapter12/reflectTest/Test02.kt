package com.easykotlin.test.chapter12.reflectTest

fun test01() {
    val digitRegex = "\\d+".toRegex()
    digitRegex.matches("7").apply {
        println(this)
    }
    digitRegex.matches("6").also {
        println(it)
    }

    digitRegex.matches("5").run {
        println(this)
    }

    println(digitRegex.matches("x"))
}

fun test02() {
    val digitRegex = "\\d+".toRegex()
    val isDigit = digitRegex::matches
    println(isDigit("7"))
    println(isDigit("8"))
    println(isDigit("44"))
    println(isDigit("x"))
}

fun main(args: Array<String>) {
    test01()
    test02()
}