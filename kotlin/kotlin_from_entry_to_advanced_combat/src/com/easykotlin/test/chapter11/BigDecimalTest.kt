package com.easykotlin.test.chapter11

import java.math.BigDecimal

operator fun BigDecimal.inc() = this + BigDecimal.ONE
operator fun BigDecimal.dec() = this - BigDecimal.ONE

fun BigDecimalIncTest() {
    var bigDecimal1 = BigDecimal(100)
    var bigDecimal2 = BigDecimal(100)

    val tmp1 = bigDecimal1 ++
    val tmp2 = ++ bigDecimal2

    println(tmp1)
    println(tmp2)

    println(bigDecimal1)
    println(bigDecimal2)
}

fun BigDecimalDecTest() {
    var bigDecimal1 = BigDecimal(100)
    var bigDecimal2 = BigDecimal(100)

    val tmp1 = bigDecimal1 --
    val tmp2 = -- bigDecimal2

    println(tmp1)
    println(tmp2)
    println(bigDecimal1)
    println(bigDecimal2)
}

fun main(args: Array<String>) {
    println("自增测试：")
    BigDecimalIncTest()

    println("自减测试：")
    BigDecimalDecTest()
}