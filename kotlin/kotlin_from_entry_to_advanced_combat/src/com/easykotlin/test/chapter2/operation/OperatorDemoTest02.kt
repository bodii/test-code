package com.easykotlin.test.chapter2.operation

import org.junit.Test
import org.junit.runner.RunWith
import org.junit.runners.JUnit4


data class Counter(var index: Int)  // 声明一个计数类
operator fun Counter.plus(incerment: Int): Counter { // 重载“+” 运算符
    return Counter(index + incerment)  // 计数器的实现: index 加上increment
}

@RunWith(JUnit4::class)
class OperatorDemoTest02 {

    @Test
    fun testCounterIndexPlus() {
        val c = Counter(1)
        val cplus = c + 10
        println(cplus)
    }
}