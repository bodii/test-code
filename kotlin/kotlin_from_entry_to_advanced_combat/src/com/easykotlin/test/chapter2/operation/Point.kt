package com.easykotlin.test.chapter2.operation

data class Point(val x: Int, val y: Int) // 声明数据类Point
operator fun Point.unaryMinus() = Point(-x, -y) // operator 修饰符修饰一个重载操作符函数