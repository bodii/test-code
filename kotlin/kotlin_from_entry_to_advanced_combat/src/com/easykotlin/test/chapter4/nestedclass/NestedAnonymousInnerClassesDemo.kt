package com.easykotlin.test.chapter4.nestedclass

class NestedAnonymousInnerClassesDemo {
    class AnonymousInnerClassDemo {
        var isRunning = false
        fun doRun() {
            Thread(object: Runnable {
                override fun run() {
                    isRunning = true
                    println("doRun: i am running, isRunning = $isRunning")
                }
            })
        }
    }
}