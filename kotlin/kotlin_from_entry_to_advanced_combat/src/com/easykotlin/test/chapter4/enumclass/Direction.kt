package com.easykotlin.test.chapter4.enumclass

/**
 * 使用enum class 声明一个Direction枚举类型
 * 相比于字符串常量，使用枚举能够实现类型安全。枚举类有两个内置的属性：
 *  public final val name: String
 *  public final val ordinal: Int
 *
 *  每一个枚举都是枚举类的实例，它们可以被初始化
 */
enum class Direction {
    // 每个枚举常量都是一个对象，用逗号分隔
    NORTH, SOUTH, WEST, EAST
}

fun main(args: Array<String>) {
    val north = Direction.NORTH // 访问枚举中的NORTH对象
    println(north.name)

    println(north.ordinal) // ordinal 属性

    println(north is Direction) // north的类型是Direction
}