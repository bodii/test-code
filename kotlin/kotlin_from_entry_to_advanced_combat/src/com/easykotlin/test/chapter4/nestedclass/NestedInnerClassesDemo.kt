package com.easykotlin.test.chapter4.nestedclass

/**
 * 如果一个类inner想要访问外部类Outer中的成员，可以在这类前面
 * 添加修饰符inner。内部类会带有一个对外部类的对象引用
 * */
class NestedInnerClassesDemo {
    class Outer {
        private val zero: Int = 0
        val one: Int = 1

        // 使用Inner关键字声明内部类
        inner class Inner {
            fun accessOuter() {
                println(zero)
                println(one)
            }
        }
    }
}

fun main(args: Array<String>) {
    val innerClass = NestedInnerClassesDemo.Outer().Inner().accessOuter()
}