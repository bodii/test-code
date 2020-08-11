package com.easykotlin.test.chapter12

/**
 * 如果注解用作另一个注解的参数时，则其名称不能以@字符为前缀。
 */
annotation class AnnotX(val value: String)

annotation class AnnotY(
        val message: String,
        val annotX: AnnotX = AnnotX("X"))