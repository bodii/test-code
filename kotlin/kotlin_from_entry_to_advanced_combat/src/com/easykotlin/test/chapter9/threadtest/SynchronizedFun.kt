package com.easykotlin.test.chapter9.threadtest

import java.io.File

fun appendFileSync(text: String, destFile: String) {
    val f = File(destFile)
    if (!f.exists())
        f.createNewFile()

    synchronized(Thread.currentThread()) {
        f.appendText(text)
    }
}

fun main(args: Array<String>) {
    appendFileSync("test", "src/com/easykotlin/test/chapter9/threadtest/demo.txt")
}