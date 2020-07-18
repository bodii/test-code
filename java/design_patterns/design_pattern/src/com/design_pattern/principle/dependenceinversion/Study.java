package com.design_pattern.principle.dependenceinversion;

public class Study {
	public void studyCourse(ICourse course) {
		course.studyCourse();
	}
 
}
