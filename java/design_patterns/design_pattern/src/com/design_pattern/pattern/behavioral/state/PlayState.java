package com.design_pattern.pattern.behavioral.state;

public class PlayState extends CourseVideoState {
    @Override
    public void play() {
        System.out.println("正常播放视频");
    }

    @Override
    public void speed() {
        System.out.println("快进");
    }

    @Override
    public void pause() {
        System.out.println("暂停");
    }

    @Override
    public void stop() {
        System.out.println("停止");
    }
}
