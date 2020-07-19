package com.design_pattern.pattern.creational.builder.v2;

public class Test {
    public static void main(String[] args) {
        Course course = 
            new Course.CourseBuilder().builderCourseName("课程名称")
            .builderCoursePPT("课程PPT").builderCourseVideo("课程视频")
            .builderCourseArticle("课程文章").builderCourseQA("课程问答")
            .build();
        
        System.out.println(course);
    }
}