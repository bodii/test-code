package com.easykotlin.test.chapter9

import java.io.BufferedReader
import java.io.InputStreamReader


/**
 * 给String扩展execute()函数
 */
fun String.execute(): Process {
    val runtime = Runtime.getRuntime()
    return runtime.exec(this)
}

fun Process.text(): String {
    var output = ""
    // 输出shell执行的结果
    val inputStream = this.inputStream
    val isr = InputStreamReader(inputStream)
    val reader = BufferedReader(isr)
    var line: String? = ""
    while (line != null) {
        line = reader.readLine()
        output += line + "\n"
    }

    return output
}


fun main(args: Array<String>) {
    val p = "ls -al".execute()

    val exitCode = p.waitFor()
    val text = p.text()

    println(exitCode)
    println(text)
}