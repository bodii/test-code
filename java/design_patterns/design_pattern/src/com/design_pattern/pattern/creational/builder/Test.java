package com.design_pattern.pattern.creational.builder;

public class Test {
    public static void main(String[] args) {
        CourseBuilder courseBuilder = new CourseActualBuilder();
        Coach coach = new Coach();
        coach.setCourseBuilder(courseBuilder);
        Course course = coach.makeCourse("课程名称", "课程PPT", "课程视频", "课程文章", "课程问答");

        System.out.println(course);
    }
}