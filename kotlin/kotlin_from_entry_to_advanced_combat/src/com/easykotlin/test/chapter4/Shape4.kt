package com.easykotlin.test.chapter4

abstract class Shape4 {
    abstract var width: Double
    abstract var height: Double
    abstract var radius: Double

    abstract fun area(): Double

    /**
     * onClick() default final 不可被覆盖重写
     * 如果想要开放给子类重新实现这个函数，可以在前面加上open关键字
     *
     */
    open fun onCheck() {
        println("I am Clicked!")
    }
}

class Rectangle4(override var width: Double, override var height: Double,
                 override var radius: Double): Shape4() {

    override fun area(): Double {
        return width * height
    }

    override fun onCheck() {
        println("${this::class.simpleName} is clicked!")
    }
}


fun main(args: Array<String>) {
    var r = Rectangle4(3.0, 4.0, 0.0)
    println(r.area())
    r.onCheck()
}