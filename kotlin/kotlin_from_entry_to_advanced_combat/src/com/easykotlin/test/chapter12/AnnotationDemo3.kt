package com.easykotlin.test.chapter12

import org.junit.Test
import org.junit.runner.RunWith
import org.junit.runners.JUnit4

@RunWith(JUnit4::class)
class AnnotationDemo3 {

    @Test
    fun testAnno() {
        val sword = SwordTest()
        sword.testCase("10000")
    }
}