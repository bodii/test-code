package com.easykotlin.test.chapter11

import kotlin.math.sqrt

class Point2(val x: Int, val y: Int) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as Point2
        if (x != other.x) return false
        if (y != other.x) return false

        return true
    }

    override fun hashCode(): Int {
        var result = x
        result = 31 * result + y

        return result
    }

    operator fun compareTo(other: Point2): Int {
        val thisNorm = sqrt((x * x + y * y).toDouble())
        val otherNorm = sqrt((other.x * other.x + other.y * other.y).toDouble())

        return thisNorm.compareTo(otherNorm)
    }
}

fun main(args: Array<String>) {
    val p1 = Point2(1, 1)
    val p2 = Point2(1, 1)
    val p3 = Point2(1, 3)

    println(p1 >= p2)
    println(p3 > p1)
}