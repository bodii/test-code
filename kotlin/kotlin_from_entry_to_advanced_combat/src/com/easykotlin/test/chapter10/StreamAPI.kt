package com.easykotlin.test.chapter10

import java.io.*

fun String.stream() = FileInputStream(this)
fun FileInputStream.buffered() = BufferedInputStream(this)
fun InputStream.reader(charset: String) = InputStreamReader(this)
fun Reader.readLines(): List<String> {
    val result = arrayListOf<String>()
    forEachLine {
        result.add(it)
    }

    return result
}

fun main(args: Array<String>) {
    val lines =
            "src/com/easykotlin/test/chapter10/data.txt"
                    .stream()
                    .buffered()
                    .reader("utf-8")
                    .readLines()
    lines.forEach(::println)
}