package com.design_pattern.pattern.creational.builder;

public class Coach {
    private CourseBuilder CourseBuilder;

    public void setCourseBuilder(CourseBuilder CourseBuilder) {
        this.CourseBuilder = CourseBuilder;
    }

    public Course makeCourse(String courseName, String coursePPT, 
            String courseVideo, String courseArticle, String courseQA) {
        CourseBuilder.builderCourseName(courseName);
        CourseBuilder.builderCoursePPT(coursePPT);
        CourseBuilder.builderCourseVideo(courseVideo);
        CourseBuilder.builderCourseArticle(courseArticle);
        CourseBuilder.builderCourseQA(courseQA);

        return CourseBuilder.makeCourse();
    }

}