package com.design_pattern.principle.demeter;

import java.util.List;
import java.util.ArrayList;

public class TeamLeader  {
    public void checkCourseNumber() {
        List<ICourse> courseList = new ArrayList<>();
        for (int i = 0; i < 20; i ++)
            courseList.add(new ICourse(){});

        System.out.println("size: " + courseList.size());
    }
}