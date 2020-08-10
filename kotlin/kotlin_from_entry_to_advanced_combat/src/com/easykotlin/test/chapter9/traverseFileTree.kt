package com.easykotlin.test.chapter9

import java.io.File

fun traverseFileTree(filename: String) {
    val f = File(filename)
    val fileTreeWalk = f.walk()
    fileTreeWalk.iterator().forEach {
        println(it.absolutePath)
        println(it.name)
    }
}

fun main(args: Array<String>) {
    traverseFileTree("src/com/easykotlin/test/chapter9/")
}