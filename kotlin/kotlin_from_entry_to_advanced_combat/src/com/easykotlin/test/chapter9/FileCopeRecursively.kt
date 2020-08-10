package com.easykotlin.test.chapter9

import java.io.File

fun fileCopy(path: String, targetPath: String, overwrite: Boolean = true) {
    val f = File(path)
    var target = File(targetPath)
    f.copyRecursively(target)
}

fun main(args: Array<String>) {
    val path = "src/com/easykotlin/test/chapter9/"
    println(fileCopy(path, path + "/copy/"))
}