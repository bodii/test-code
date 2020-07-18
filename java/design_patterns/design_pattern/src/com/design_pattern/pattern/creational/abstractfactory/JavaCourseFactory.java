package com.design_pattern.pattern.creational.abstractfactory;

public class JavaCourseFactory implements CourseFactory {

    @Override
    public Video getVideo() {
       return new JavaVideo();
    }

    @Override
    public Article getArticel() {
        return new JavaAtricle();
    }
    
}