package com.easykotlin.test.chapter9

fun main(args: Array<String>) {
    val re = Regex("[0-9]+")
    val p = re.toPattern()
    val m = p.matcher("888ABC999")
    while (m.find())
        println(m.group())
}