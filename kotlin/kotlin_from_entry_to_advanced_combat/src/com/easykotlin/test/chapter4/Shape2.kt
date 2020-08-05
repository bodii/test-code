package com.easykotlin.test.chapter4

abstract class Shape2 {
    abstract var width: Double
    abstract var height: Double
    abstract var radius: Double
    abstract fun area(): Double
}

class Rectangle2(override var width: Double, override var height: Double,
            override var radius: Double): Shape2() {

    override fun area(): Double {
        return height * width
    }
}

class Circle2(override var width: Double, override var height: Double,
              override var radius: Double): Shape2() {

    override fun area(): Double {
        return 3.14 * radius * radius
    }

}


fun main(args: Array<String>) {
    val r = Rectangle2(3.0, 4.0, 0.0)
    println(r.area())

    val c = Circle2(0.0, 0.0, 4.0)
    println(c.area())
}