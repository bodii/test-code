package com.easykotlin.test.chapter7

// 过滤函数
data class Student(var id: Long, val name: String, var age: Int, var score: Int) {
    override fun toString(): String {
        return "Student(id=$id, name=$name, age=$age, score=$score)"
    }
}

fun main(args: Array<String>) {
    val studentList = listOf(
            Student(1, "Jack", 18, 90),
            Student(2, "Rose", 17, 80),
            Student(3, "Alice", 16, 70)
    )

    // 过滤出年龄大于等于18行岁的学生
    println(studentList.filter{ it.age >= 18 })

    // 过滤出分数小于80分的学生
    println(studentList.filter { it.score < 80 })

    // 如果想要通过访问下标来过滤，可以使用filterIndexed()函数
    val list = listOf(1,2,3,4,5,6,7)
    val newList = list.filterIndexed {
        index, it ->
        index % 2 == 0 && it > 3
    }
    println(newList)
}