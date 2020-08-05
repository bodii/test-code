package com.easykotlin.test.chapter4.nestedclass

fun doStop() {
    var isRunning = true
/*    Thread({
        isRunning = false
        println("doStop: i am not running, isRunning = $isRunning")
    }).start()*/

    Thread {
        isRunning = false
        println("doStop: i am not running, isRunning = $isRunning")
    }.start()
}

fun doWait() {
    var isRunnig = true
    val wait = Runnable {
        isRunnig = false
        println("doWait: i am waiting, isRunning = $isRunnig")
    }
    Thread(wait).start()
}

fun doNotify() {
    var isRunnig = true
    val wait = {
        isRunnig = false
        println("doNotify: i notify, isRunning = $isRunnig")
    }
    Thread(wait).start()
}

fun main(args: Array<String>) {

    doWait()
    doNotify()
    doStop()
}