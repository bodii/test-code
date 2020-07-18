package com.design_pattern.principle.dependenceinversion;

public class Study3 {
    private ICourse course;

    public void setCourse(ICourse course) {
        this.course = course;
    }

    public void studyCourse() {
        course.studyCourse();
    }
}