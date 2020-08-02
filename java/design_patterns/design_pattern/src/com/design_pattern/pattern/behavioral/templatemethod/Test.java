package com.design_pattern.pattern.behavioral.templatemethod;

public class Test {
    public static void main(String[] args) {
        System.out.println("后端设计模式课程start----");
        ACourse designPatternCourse = new DesignPtternCourse(true);
        designPatternCourse.makeCourse();
        System.out.println("后端设计模式课程end ----");

        System.out.println();

        System.out.println("前端课程start----");
        ACourse feCourse = new FECourse();
        feCourse.makeCourse();
        System.out.println("前端设计模式课程end----" );
    }
}
