package com.easykotlin.test.chapter9

import java.net.URL

fun getUrlBytes(url: String): ByteArray {
    return URL(url).readBytes()
}

fun main(args: Array<String>) {
    println(getUrlBytes("http://cn.bing.com"))
}