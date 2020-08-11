package com.easykotlin.test.chapter12

import java.util.*
import kotlin.reflect.full.declaredFunctions

fun testAnnotProcessing() {
    val sword = SwordTest()
    val kClass = sword::class
//    val kClass: KClass<out SwordTest> = sword::class
    val declaredFunctions = kClass.declaredFunctions
    println(declaredFunctions)

    for (f in declaredFunctions) {
        // 处理TestCase注解，使用其中的元数据
        f.annotations.forEach {
            if (it is TestCase) {
                val id = it.id
                // 注解处理逻辑
                 doSomething(id)
                // 如果想通过反射来调用函数，可以直接使用call()函数
                f.call(sword, id)
                // 等价于: f.javaMethod?.invoke(sword, id)
            }
        }
    }
}

private fun doSomething(id: String) {
    println("Do Something in Annotation Processing $id ${Date()}")
}

fun main(args: Array<String>) {
    testAnnotProcessing()
}