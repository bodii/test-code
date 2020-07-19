package com.design_pattern.pattern.creational.builder.v2;

public class Course {
    private String courseName;
    private String coursePPT;
    private String courseVideo;
    private String courseArticle;
    private String courseQA;

    public static class CourseBuilder {
        private String courseName;
        private String coursePPT;
        private String courseVideo;
        private String courseArticle;
        private String courseQA;

        public CourseBuilder builderCourseName(String courseName) {
            this.courseName = courseName;
            return this;
        }

        public CourseBuilder builderCoursePPT(String coursePPT) {
            this.coursePPT = coursePPT;
            return this;
        }

        public CourseBuilder builderCourseVideo(String courseVideo) {
            this.courseVideo = courseVideo;
            return this;
        }

        public CourseBuilder builderCourseArticle(String courseArticle) {
            this.courseArticle = courseArticle;
            return this;
        }

        public CourseBuilder builderCourseQA(String courseQA) {
            this.courseQA = courseQA;
            return this;
        }

        public Course build() {
            return new Course(this);
        }
    }

    public Course(CourseBuilder courseBuilder) {
        this.courseName = courseBuilder.courseName;
        this.coursePPT = courseBuilder.coursePPT;
        this.courseVideo = courseBuilder.courseVideo;
        this.courseArticle = courseBuilder.courseArticle;
        this.courseQA = courseBuilder.courseQA;
    }
    
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