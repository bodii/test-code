package com.design_pattern.pattern.behavioral.command;

public class CourseVideo {
    private String name;

    public CourseVideo(String name) {
        this.name = name;
    }

    public void open() {
        System.out.println(name + "课程视频被打开");
    }

    public void close() {
        System.out.println(name + "课程视频被关闭");
    }
}
