package com.easykotlin.test.chapter4.`interface`

class ProjectMilestoneServiceImpl(override val name: String, override val owner: String) : ProjectService, MilestoneService {
    override fun save(project: Project) {
        TODO("Not yet implemented")
    }

    override fun print() {
//        super.print() // error 编译器无法知道我们想要调用的是哪个print函数，这个叫覆盖冲突
        super<ProjectService>.print() // 调用ProjectService接口中的print()函数
        super<MilestoneService>.print() // 使用MilestoneService接口中的print()函数

    }
}

fun main(args: Array<String>) {
    val projectMilestoneServiceImpl = ProjectMilestoneServiceImpl("aa", "bb");
    projectMilestoneServiceImpl.print()
}