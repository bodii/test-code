package com.easykotlin.test.chapter12.reflectTest

import java.lang.reflect.ParameterizedType

/**
 * Kotlin中的反射获取泛型代码
 */

class A<T>

open class C<T>

class B<T>: C<Int>() // 继承父类

fun fooA() {
    val parameterizedType = A<Int>()::class.java.genericSuperclass as ParameterizedType
    val actualTypeArguments = parameterizedType.actualTypeArguments
    for (type in actualTypeArguments) {
        val typeName = type.typeName
        println("typeName = ${typeName}")
    }
}

fun fooB() {
    val parameterizedType = B<Int>()::class.java.genericSuperclass as ParameterizedType
    val actualTypeArguments = parameterizedType.actualTypeArguments
    for (type in actualTypeArguments) {
        val typeName = type.typeName
        println("typeName = $typeName")
    }
}


fun main(args: Array<String>) {
//    fooA()
    fooB()
}