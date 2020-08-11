package com.easykotlin.test.chapter11

class Point(val x: Int, val y: Int) {
    operator fun unaryMinus() = Point(-x, -y)
    override fun toString(): String {
        return "Point(x=$x, y=$y)"
    }
}

fun main(args: Array<String>) {
    val p1 = Point(1, 1)
    println(-p1)
}