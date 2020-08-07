package com.easykotlin.test.chapter5

fun testLetFun() {
    1.let { println(it) }       // 1
    "ABC".let { println(it) }   // ABC
    // run myfun()
    myfun().let {
        print(it)
    }
}

fun main(args: Array<String>) {
    testLetFun()
}