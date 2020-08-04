package com.easykotlin.test.chapter4

class AnEmptyClass // 一个空类

fun main(args: Array<String>) {
    val anEmptyClass = AnEmptyClass()  // Kotlin中不需要使用new
    println(anEmptyClass)
    println(anEmptyClass is AnEmptyClass)
    println(anEmptyClass::class)
}