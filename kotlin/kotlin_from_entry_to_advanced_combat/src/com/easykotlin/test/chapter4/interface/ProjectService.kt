package com.easykotlin.test.chapter4.`interface`

interface ProjectService {
    val name: String
    val owner: String
    fun save(project: Project)

    fun print() {
        println("I am project")
    }
}
