package com.easykotlin.test.chapter3.typing

fun main(args: Array<String>) {
    println(strLength(null))
    println(strLength("abc"))

//    val str: String = null  // error：String类型不能是null
    val str: String? = null
    println(str)

    nullTypeTest()

    nullTypeTest002()
}

fun strLength(s: String?): Int {
    return s?.length ?: 0
}

fun nullTypeTest() {
    println("null == null: " + (null == null))
    println("null != null: " + (null != null))
    println("null is Any: " + (null is Any))
    println("null is Any?: " + (null is Any?))
}

fun nullTypeTest002() {
    var a = null
    println("a: " + a)
//    a = 1
//    println("a: " + a)

    var nullableStr: String? = null
    println("nullableStr: " + nullableStr)
    println("nullableStr length: " + nullableStr?.length)

    var ableStr: String? = "abc"
    println(ableStr?.length)

    nullableStr = null
//    println("nullableStr!! : " + nullableStr!!.length)
    var s = nullableStr ?: "NULL"
    println("s: " + s)
}
