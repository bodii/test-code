package com.easykotlin.test.chapter9

import java.io.File
import java.nio.charset.Charset

fun writeFile(text: String, destFile: String) {
    val f = File(destFile)
    if (!f.exists())
        f.createNewFile()

    f.writeText(text, Charset.defaultCharset())
}

fun main(args: Array<String>) {
    writeFile("abcdef", "src/com/easykotlin/test/chapter9/exampleFile.txt")
}