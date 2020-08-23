package com.design_pattern.pattern.behavioral.state;

public class Test {
    public static void main(String[] args) {
        CourseVideoContext courseVideoContext = new CourseVideoContext();
        courseVideoContext.setCourseVideoState(new PlayState());
        System.out.println("当前状态: " + courseVideoContext.getCourseVideoState().getClass().getSimpleName());
        
        courseVideoContext.setCourseVideoState(new PauseState());
        System.out.println("当前状态: " + courseVideoContext.getCourseVideoState().getClass().getSimpleName());

        courseVideoContext.setCourseVideoState(new SpeedState());
        System.out.println("当前状态: " + courseVideoContext.getCourseVideoState().getClass().getSimpleName());

        courseVideoContext.setCourseVideoState(new StopState());
        System.out.println("当前状态: " + courseVideoContext.getCourseVideoState().getClass().getSimpleName());


    }
    
}