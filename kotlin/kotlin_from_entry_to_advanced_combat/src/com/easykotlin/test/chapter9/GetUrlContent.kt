package com.easykotlin.test.chapter9

import java.net.URL
import java.nio.charset.Charset

fun getUrlContent(url: String): String {
    return URL(url).readText(Charset.defaultCharset())
}

fun main(args: Array<String>) {
    val bingText = getUrlContent("http://cn.bing.com")
    println(bingText)
}