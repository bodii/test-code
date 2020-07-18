package com.design_pattern.principle.dependenceinversion;

public class PHPCourse  implements ICourse{
	@Override
	public void studyCourse() {
		System.out.println("我正在学习PHP课程");
	}
}
