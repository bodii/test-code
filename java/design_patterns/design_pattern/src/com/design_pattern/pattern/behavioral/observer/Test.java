package com.design_pattern.pattern.behavioral.observer;

public class Test {
    public static void main(String[] args) {
        Course course = new Course("Java设计课程");
        Question question = new Question();
        question.setUserName("Jack");
        question.setQuestionContent("Java的程序如何编写？");

        Teacher tom_teacher = new Teacher("Tom");
        Teacher lous_teacher = new Teacher("Lous");

        course.addObserver(tom_teacher);
        course.addObserver(lous_teacher);
        course.produceQuestion(course, question);
    }
}
