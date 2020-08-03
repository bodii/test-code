package com.easykotlin.test.chapter2

import java.util.Date
import java.text.SimpleDateFormat

fun main(args: Array<String>) {
    println("Hello world!");
    println(SimpleDateFormat("Y-MM-dd HH:mm:ss").format(Date()))
}