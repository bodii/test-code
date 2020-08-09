package com.easykotlin.test.chapter7

fun main(args: Array<String>) {
    // 创建不可变List
    val list = listOf(1, 2, 3, 4, 5, 6, 7)
    // 创建可变MutableList
    val mutableList = mutableListOf(1, 2, 3, 4, 5, 6, 7)

    // 创建不可变Set
    val set = setOf(1, 2, 3, 4, 5, 6, 7)
    // 创建可变MutableSet
    val mutableSet = mutableSetOf("a", "b", "C")

    // 创建不可变Map
    val map = mapOf(1 to "a", 2 to "b", 3 to "c")
    // 创建可变MutableMap
    val mutableMap = mutableMapOf(1 to "x", 2 to "y", 3 to "z");

    // 如果创建没有元素的空List，使用listOf()即可。不过这个时候，变量的类型不能省略
    // 因为这里的funlistOf():List泛型参数T编译器无法推断出来。setOf()、mapOf()同理。
    // 显式声明List的元素类型为Int
    val emptyList: List<Int> = listOf()
    // 显式声明Set的元素类型为Int
    val emptySet: Set<Int> = setOf()
    // 显式声明Map的元素类型为Int,String键值对
    val emptyMap: Map<Int, String> = mapOf()

}