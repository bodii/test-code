package com.easykotlin.test.chapter11

class Int2(var v: Int) {
    operator fun plusAssign(b: Int2) {
        v = v + b.v
    }

    operator fun minusAssign(b: Int2) {
        v = v - b.v
    }

    operator fun timesAssign(b: Int2) {
        v = v * b.v
    }

    operator fun divAssign(b: Int2) {
        v = v / b.v
    }

    operator fun remAssign(b: Int2) {
        v = v % b.v
    }

    override fun toString(): String {
        return "$v"
    }
}


fun main(args: Array<String>) {
    var a = Int2(10)
    val b = Int2(2)
    a += b
    println(a)

    a -= b
    println(a)

    a /= b
    println(a)

    a *= b
    println(a)

    a %= b
    println(a)
}
