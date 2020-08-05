package com.easykotlin.test.chapter4.nestedclass

/*
    普通嵌套类没有持有外部类的引用，所以是无法访问外部变量的
 */
class NestedClassesDemo {
    class Outer {
        private val zero: Int = 0
        val one: Int = 1

        class Nested {
            fun getTwo() = 2
            fun accessOuter() = {
//                println(zero) // error: cannot access outer class
//                println(one)  // error: cannot access outer class
            }
        }
    }
}