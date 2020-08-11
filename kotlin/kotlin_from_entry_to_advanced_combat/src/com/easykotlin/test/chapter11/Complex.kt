package com.easykotlin.test.chapter11

class Complex {
    var real: Int = 0
    var image: Int = 0

    constructor()
    constructor(real: Int, image: Int) {
        this.real = real
        this.image = image
    }

    operator fun plus(c: Complex): Complex {
        return Complex(real + c.real, image + c.image)
    }

    operator fun minus(c: Complex): Complex {
        return Complex(real - c.real, image - c.image)
    }

    operator fun times(c: Complex): Complex {
        return Complex(real * c.real - image * c.image,
                c.image + image * c.real)
    }

    override fun toString(): String {
        val img = if (image >= 0) "+ ${image}i" else "${image}i"
        return "$real $img"
    }

}


fun main(args: Array<String>) {
    val c1 = Complex(1, 1)
    val c2 = Complex(2, 2)
    val p = c1 + c2
    val m = c1 - c2
    val t = c1 * c2

    println(p)
    println(m)
    println(t)
}


