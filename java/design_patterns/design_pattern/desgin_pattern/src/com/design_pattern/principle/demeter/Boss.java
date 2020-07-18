package com.design_pattern.principle.demeter;

public class Boss {
    public void commandcheckCourseNumber(TeamLeader teamLeader) {
        teamLeader.checkCourseNumber();
    }
}