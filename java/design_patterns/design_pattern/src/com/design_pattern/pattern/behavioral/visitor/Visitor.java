package com.design_pattern.pattern.behavioral.visitor;

public class Visitor implements Ivisitor {
    @Override
    public void visit(Course course) {
        if (course instanceof CodingCourse)
            System.out.println("实战课程：" + course.getName() + " 价格：" + ((CodingCourse) course).getPrice());
        else if (course instanceof FreeCourse)
            System.out.println("免费课程：" + course.getName());
    }
}
