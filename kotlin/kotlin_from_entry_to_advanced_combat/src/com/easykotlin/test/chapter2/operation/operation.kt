package com.easykotlin.test.chapter2.operation

fun operation01() {
    val a = 10
    val b = 3
    println(a + b)

    println(a - b)

    println(a / b)

    println(a % b)

    println(a..b)

    println(b.rangeTo(a))
}

fun operation02() {
    println("" + 1) // String类型重载了加法操作符

//    println(1 + "") // error Int类型没有重载操作符plus(other:String)
    println(1.toString() + "")
}


fun main(args: Array<String>) {
    operation01()
    operation02()
}

