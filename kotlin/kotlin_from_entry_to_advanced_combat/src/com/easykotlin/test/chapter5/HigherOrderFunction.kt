package com.easykotlin.test.chapter5

fun main(args: Array<String>) {
    val strList = listOf("a", "ab", "abc", "abcd", "abcde", "abcdef")
    
    val f = fun (x: Int) = x % 2 == 1  // 判断输入的Int是否奇数
    val g = fun (s: String) = s.length  // 返回输入的字符串参数的长度

//    val h = fun (g: (String) -> Int, f: (Int) -> Boolean): (String) -> Boolean {
//        return { f(g(it)) }
//    }

    val h = fun(g: G, f: F): H {
        return { f(g(it)) }
    }

    println(strList.filter(h(g, f)))
}



typealias G = (String) -> Int
typealias F = (Int) -> Boolean
typealias H = (String) -> Boolean