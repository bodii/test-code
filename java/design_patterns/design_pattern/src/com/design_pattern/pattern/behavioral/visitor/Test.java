package com.design_pattern.pattern.behavioral.visitor;

import java.util.ArrayList;
import java.util.List;

public class Test {
    public static void main(String[] args) {
        List<Course> courseList = new ArrayList<>();
        FreeCourse freeCourse = new FreeCourse();
        freeCourse.setName("Spring MVC数据绑定");

        CodingCourse codingCourse = new CodingCourse();
        codingCourse.setName("Java 设计模式精讲");
        codingCourse.setPrice(299);

        courseList.add(freeCourse);
        courseList.add(codingCourse);

        for (Course course : courseList)
            course.accept(new Visitor());

    }
}
