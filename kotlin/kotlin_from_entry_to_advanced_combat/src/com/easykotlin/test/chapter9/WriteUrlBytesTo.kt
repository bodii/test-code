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

    writeUrlBytesTo(path + "OIP.dCFZwNZRvKc89JHiYNC4fwHaE8.jpg", "https://tse2-mm.cn.bing.net/th/id/OIP.dCFZwNZRvKc89JHiYNC4fwHaE8?pid=Api&rs=1")
}