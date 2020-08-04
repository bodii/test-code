package com.easykotlin.test.chapter2.operation

import org.junit.Test
import org.junit.runner.RunWith
import org.junit.runners.JUnit4

@RunWith(JUnit4::class)
class OperatorDemoTest {

    @Test
    fun testPointUnaryMinus() {
        val p = Point(1, 1)
        val np = -p // 直接使用unaryMinus()重载函数操作符"-"
        println(np)
    }
}