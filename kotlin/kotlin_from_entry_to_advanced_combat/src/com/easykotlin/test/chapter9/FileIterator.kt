package com.easykotlin.test.chapter9

import java.io.File

/**
 * 遍历当前文件下所有的子目录文件，将其存入一个Iterator中
 */
fun getFileIterator(filename: String): Iterator<File> {
    val f = File(filename)
    val fileTreeWalk = f.walk()
    return fileTreeWalk.iterator()
}

fun getFileSequenceBy(filename: String, p: (File) -> Boolean): Sequence<File> {
    val f = File(filename)
    return f.walk().filter(p)
}


fun main(args: Array<String>) {
    val filePath = "src/com/easykotlin/test/chapter9/"
    val fileTreeIter = getFileIterator(filePath)

    val fileSequence = getFileSequenceBy(filePath) {
        it.name.endsWith(".kt")
    }

    fileSequence.iterator().forEach {
        println(it.name)
    }

}