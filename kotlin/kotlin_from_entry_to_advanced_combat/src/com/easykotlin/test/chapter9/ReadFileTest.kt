package com.easykotlin.test.chapter9

import java.io.File
import java.nio.charset.Charset

fun getFileContent(filename: String): String {
    val f = File(filename)
    return f.readText(Charset.forName("UTF-8"))
}

fun getFileLines(filename: String): List<String> {
    return File(filename).readLines(Charset.forName("UTF-8"))
}

fun getFileBytes(filename: String): ByteArray {
    val f = File(filename)
    println(f.readBytes().joinToString(separator = " "))
    return f.readBytes()
}

fun getBufferedReader(filename: String): String {
    return File(filename).bufferedReader(Charset.forName("UTF-8")).readText()
}

fun main(args: Array<String>) {
    val filePath = "./src/com/easykotlin/test/chapter9/exampleFile.txt"
    val f = File(filePath)
//    f.writeText("abcdefg")
    println(getFileContent(filePath))

    println(getFileLines(filePath))

    println(getFileBytes(filePath))

    println(getBufferedReader(filePath))

}