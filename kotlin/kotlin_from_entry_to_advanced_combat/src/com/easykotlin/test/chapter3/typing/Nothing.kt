package com.easykotlin.test.chapter3.typing

// Nothing类型的表达式计算结果是永远不会返回的（与Java中void相同）
// Nothing?可以包含一个值null
// Nothing?唯一允许的值是null,可被用作任何可空类型的空引用
fun test003() {
    var nul: Nothing? = null
//    nul = 1  // error
//    nul = true // error
    nul = null
    println(nul)
}

fun main(args: Array<String>) {
    test003()
}