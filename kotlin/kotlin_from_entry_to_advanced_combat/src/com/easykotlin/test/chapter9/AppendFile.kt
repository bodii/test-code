package com.easykotlin.test.chapter9

import java.io.File
import java.nio.charset.Charset

fun appendFile(text: String, destFile: String) {
    val f = File(destFile)
    if (!f.exists())
        f.createNewFile()

    f.appendText(text, Charset.defaultCharset())
}

fun appendBytes(bytes: ByteArray, destFile: String) {
    val f = File(destFile)
    if (!f.exists())
        f.createNewFile()

    f.appendBytes(bytes)
}

fun main(args: Array<String>) {
    val filePath = "src/com/easykotlin/test/chapter9/exampleFile.txt"
    appendFile("\nxyz", filePath)
}