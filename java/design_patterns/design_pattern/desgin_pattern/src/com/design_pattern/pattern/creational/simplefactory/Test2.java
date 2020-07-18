package com.design_pattern.pattern.creational.simplefactory;

public class Test2 {
    public static void main(String[] args) {
        VideoFactory2 vFactory2 = new VideoFactory2();
        Video video = vFactory2.getVideo(JavaVideo.class);
        video.produce();
        
    }
}