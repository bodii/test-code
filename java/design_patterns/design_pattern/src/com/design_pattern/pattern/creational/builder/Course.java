package com.design_pattern.pattern.creational.builder;

public class Course {
    private String courseName;
    private String coursePPT;
    private String courseVideo;
    private String courseArticle;
    private String courseQA;
    
    public String getCourseName() {
        return courseName;
    }

    public String getCoursePPT() {
        return coursePPT;
    }

    public String getCourseVideo() {
        return courseVideo;
    }

    public String getCourseArticle() {
        return courseArticle;
    }

    public String getCourseQA() {
        return courseQA;
    }

    public void setCourseName(String courseName) {
        this.courseName = courseName;
    }

    public void setCoursePPT(String coursePPT) {
        this.coursePPT = coursePPT;
    }

    public void setCourseVideo(String courseVideo) {
        this.courseVideo = courseVideo;
    }

    public void setCourseArticle(String courseArticle) {
        this.courseArticle = courseArticle;
    }

    public void setCourseQA(String courseQA) {
        this.courseQA = courseQA;
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder();
        str.append("Course{\n");
        str.append("\tcourseName='" + courseName + "',\n");
        str.append("\tcoursePPT='" + coursePPT + "',\n");
        str.append("\tcourseVideo='" + courseVideo + "',\n");
        str.append("\tcourseArticle='" + courseArticle + "',\n");
        str.append("\tcourseQA='" + courseQA + "',\n");
        str.append("}");
        return str.toString();
    }

}