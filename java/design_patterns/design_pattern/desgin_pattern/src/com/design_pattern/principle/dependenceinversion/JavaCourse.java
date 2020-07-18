package com.design_pattern.principle.dependenceinversion;

public class JavaCourse implements ICourse {
	@Override
	public void studyCourse() {
		System.out.println("我正在学习Java课程");
	}
}
