package com.easykotlin.test.chapter9.threadtest

import java.io.File

@Synchronized fun appendFile(text: String, destFile: String) {
    val f = File(destFile)
    if (!f.exists())
        f.createNewFile()

    f.appendText(text)
}

fun main(args: Array<String>) {
    appendFile("test", "src/com/easykotlin/test/chapter9/threadtest/demo.txt")
}