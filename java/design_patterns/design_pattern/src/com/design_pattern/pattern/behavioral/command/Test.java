package com.design_pattern.pattern.behavioral.command;

public class Test {
    public static void main(String[] args) {
        CourseVideo courseVideo = new CourseVideo("Java设式精讲");
        Command openCommand = new OpenCourseVideoCommand(courseVideo);
        Command CloseCommand = new CloseCourseVideoCommand(courseVideo);

        Staff staff = new Staff();
        staff.addCommand(openCommand);
        staff.addCommand(CloseCommand);
        staff.executeCommands();
    }
}
