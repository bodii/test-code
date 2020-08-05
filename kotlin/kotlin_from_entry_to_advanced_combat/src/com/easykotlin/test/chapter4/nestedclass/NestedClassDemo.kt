package com.easykotlin.test.chapter4.nestedclass

class NestedClassDemo {
    class Outer {
        private val zero: Int = 0
        val one: Int = 1

        class Nested {
            fun getTwo() = 2
            class Nested1 {
                val three = 3
                fun getFour() = 4
            }
        }
    }
}

fun main(args: Array<String>) {
    val one = NestedClassDemo.Outer().one
    val two = NestedClassDemo.Outer.Nested().getTwo()
    val three = NestedClassDemo.Outer.Nested.Nested1().three
    val four = NestedClassDemo.Outer.Nested.Nested1().getFour()

    println("one=$one")
    println("two=$two")
    println("three=$three")
    println("four=$four")
}