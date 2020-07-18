package com.design_pattern.pattern.creational.simplefactory;

public class Test {
    public static void main(String[] args) {
        VideoFactory videoFactory = new VideoFactory();
        Video jvVideo = videoFactory.getVideo("java");
        jvVideo.produce();

        Video pvVideo = videoFactory.getVideo("python");
        pvVideo.produce();
    }
}