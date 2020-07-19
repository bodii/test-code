package com.design_pattern.pattern.creational.builder;

public abstract class CourseBuilder {
    public abstract void builderCourseName(String courseName);
    public abstract void builderCoursePPT(String coursePPT);
    public abstract void builderCourseVideo(String courseVideo);
    public abstract void builderCourseArticle(String courseArticle);
    public abstract void builderCourseQA(String courseQA);
    public abstract Course makeCourse();
}