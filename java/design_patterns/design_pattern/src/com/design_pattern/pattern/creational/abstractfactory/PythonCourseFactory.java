package com.design_pattern.pattern.creational.abstractfactory;

public class PythonCourseFactory implements CourseFactory {

    @Override
    public Video getVideo() {
        return new PythonVideo();
    }

    @Override
    public Article getArticel() {
        return new PythonArticle();
    }
    
}