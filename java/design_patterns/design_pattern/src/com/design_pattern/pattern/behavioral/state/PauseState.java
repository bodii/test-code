package com.design_pattern.pattern.behavioral.state;

public class PauseState extends CourseVideoState {

    @Override
    public void play() {
       super.courseVideoContext.setCourseVideoState(CourseVideoContext.PLAY_STATE); 
    }

    @Override
    public void speed() {
       System.out.println("暂停");

    }

    @Override
    public void pause() {
        System.out.println("暂停");

    }

    @Override
    public void stop() {
        super.courseVideoContext.setCourseVideoState(CourseVideoContext.STOP_STATE);

    }
}
