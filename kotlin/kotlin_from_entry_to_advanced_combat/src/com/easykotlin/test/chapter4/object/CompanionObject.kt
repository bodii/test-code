package com.easykotlin.test.chapter4.`object`

/**
 * Kotlin中提供了伴生对象，用companion object关键字声明
 * ! 一个类只能有一个伴生对象
 */
class DataProcessor {
    // 使用companion object 声明 Dataprocessor的伴生对象
    companion object DataProcesser {

        fun process() {
            println("I am processing data ... ")
        }
    }
}

fun main(args: Array<String>) {
    DataProcessor.process()
}