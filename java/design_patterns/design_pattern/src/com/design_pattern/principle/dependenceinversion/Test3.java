package com.design_pattern.principle.dependenceinversion;

public class Test3 {
    public static void main(String[] args) {
        Study3 study3 = new Study3();

        study3.setCourse(new JavaCourse());
        study3.studyCourse();

        study3.setCourse(new PHPCourse());
        study3.studyCourse();

        study3.setCourse(new PythonCourse());
        study3.studyCourse();
    }
}