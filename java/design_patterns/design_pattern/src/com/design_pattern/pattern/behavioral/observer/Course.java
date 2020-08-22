package com.design_pattern.pattern.behavioral.observer;

import java.util.Observable;

public class Course extends Observable {
    private String CourseName;

    public Course(String courseName) {
        CourseName = courseName;
    }

    public String getCourseName() {
        return CourseName;
    }

    public void produceQuestion(Course course, Question question) {
        System.out.println(question.getUserName() + "在" + course.CourseName + "下提了一个问题：" + question.getQuestionContent());
        setChanged();
        notifyObservers(question);
    }
}
