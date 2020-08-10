package com.easykotlin.test.chapter9

import java.io.File
import java.net.URL

fun writeUrlBytesTo(filename: String, url: String) {
    val bytes = URL(url).readBytes()
    File(filename).writeBytes(bytes)
}

fun main(args: Array<String>) {
    val path = "src/com/easykotlin/test/chapter9/"
    val url = "http://cn.bing.com"
    writeUrlBytesTo(path + "bing.txt", url)
}