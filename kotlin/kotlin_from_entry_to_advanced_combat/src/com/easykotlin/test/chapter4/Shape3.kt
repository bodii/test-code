package com.easykotlin.test.chapter4

abstract class Shape3 {
    abstract var width: Double
    abstract var height: Double
    abstract var radius: Double

    abstract fun area(): Double

    // 默认是final的，不可被覆盖重写
    fun onClick() {
        println("I am Clicked!")
    }
}

class Rectangle3(override var width: Double, override var height: Double,
                 override var radius: Double): Shape3() {

    override fun area(): Double {
        return width * height
    }
}

class Circle3(override var width: Double, override var height: Double,
              override var radius: Double): Shape3() {

    override fun area(): Double {
        return 3.14 * radius * radius
    }

}

fun main(args: Array<String>) {
    val r = Rectangle3(3.0, 4.0, 0.0)
    r.onClick()

    val c = Circle3(0.0, 0.0, 4.0)
    c.onClick()
}