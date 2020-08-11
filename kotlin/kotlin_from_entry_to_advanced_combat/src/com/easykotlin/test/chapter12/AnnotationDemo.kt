package com.easykotlin.test.chapter12

@Target(AnnotationTarget.CLASS,
        AnnotationTarget.FUNCTION,
        AnnotationTarget.VALUE_PARAMETER,
        AnnotationTarget.EXPRESSION)
@Retention(AnnotationRetention.SOURCE)
@Repeatable
@MustBeDocumented
annotation class TestCase(val id: String)

@Target(AnnotationTarget.CLASS,
        AnnotationTarget.FUNCTION,
        AnnotationTarget.VALUE_PARAMETER,
        AnnotationTarget.EXPRESSION)
@Retention(AnnotationRetention.SOURCE)
@Repeatable
@MustBeDocumented
annotation class Run

@Run
class SwordTest {

    @TestCase(id = "1")
    fun testCase(testId: String) {
        println("Run SwordTest ID = $testId")
    }
}

fun main(args: Array<String>) {
    val swordTest = SwordTest()
    swordTest.testCase("4")
}