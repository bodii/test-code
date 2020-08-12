package com.easykotlin.test.chapter12.reflectTest

import kotlin.reflect.KTypeParameter

open class BaseContainer2<T>

class Container2<T: Comparable<T>>(val elements: MutableList<T>) : BaseContainer2<Int>() {

    fun sort(): Container2<T> {
        elements.sort()
        return this
    }

    override fun toString(): String {
        return "Container(elements=$elements)"
    }
}

fun main(args: Array<String>) {
    val container2 = Container2(mutableListOf<Int>(1, 5, 7, 4, 3, 2))
    val kClass = container2::class // 获取KClass对象
    // 获取类型参数typeParameters
    val typeParameters = kClass.typeParameters
    // 这个typeParameters是一个数组，取第一个对象
    val kTypeParameter: KTypeParameter = typeParameters[0]
    // 对象kTypeParameter的属性有name, isReified，upperBounds和variance
    println(kTypeParameter.isReified)
    println(kTypeParameter.name)
    println(kTypeParameter.upperBounds)
    println(kTypeParameter.variance)

    // KClass的constructor属性中存在构造函数的信息
    val constructor = kClass.constructors
    for (KFunction in constructor) {
        KFunction.parameters.forEach {
            val name = it.name
            val type = it.type
            println("name = $name")
            println("type = $type")
            for (KTypeProjection in type.arguments) {
                println(KTypeProjection)
            }
        }
    }

}