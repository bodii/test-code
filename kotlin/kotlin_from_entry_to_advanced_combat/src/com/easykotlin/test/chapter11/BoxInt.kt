package com.easykotlin.test.chapter11

/**
 * 自定义一个BoxInt类型，然后重载times(乘法*)函数、plus(加法+) 
 *
 * */
class BoxInt(var i: Int) {
    operator fun times(x: BoxInt) = BoxInt(i * x.i)  // 使用类成员函数重载

    override fun toString(): String {
        return i.toString()
    }
}

// 使用扩展函数的方式重载
operator fun BoxInt.plus(x: BoxInt) = BoxInt(i + x.i)


fun main(args: Array<String>) {
    val a = BoxInt(3)
    val b = BoxInt(4)
    println(a + b)
    println(a * b)
}