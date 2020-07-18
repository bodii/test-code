package com.design_pattern.principle.dependenceinversion;

public class Study2 {
	private final ICourse course;
	
	public Study2(ICourse course) {
		this.course = course;
	}
	
	public void studyCourse() {
		course.studyCourse();
	}
 
}
