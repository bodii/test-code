package com.easykotlin.test.chapter2

fun main(args: Array<String>) {
    println(max(1, 2))
    println(isLeapYear(2017))
    println(isLeapYear(2020))
}

fun max(a: Int, b: Int): Int {
    // 表达式返回值
    val max = if (a > b) a else b

    return max
}

fun max3 (a: Int, b: Int): Int {
    val max = if (a > b) { // 这里{}中的内容是一个代码块
        print("Max is a")
        a                   // 最后的表达式作为该代码块的值
    } else {
        print("Max is b")
        b                   // 同上
    }

    return max
}

// error: if("a") 1
// error: if(1) println(1)
// if...else语句规则: if后的括号不能省略，括号里表达式的值必须是布尔型。
// if(true) println(1) else println(0) // if 后面的大括号可以省略
// if(true) { println(1) } else { println(0) } // 良好的编程风格是建议加上大括号

/**
 * 判断是否是闰年
 */
fun isLeapYear(year: Int): Boolean {
    var isLeapYear: Boolean // 声明一个局部变量
    if ((year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)) {
        isLeapYear = true
    } else {
        isLeapYear = false
    }
    // isLeapYear = (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)

    return isLeapYear
}