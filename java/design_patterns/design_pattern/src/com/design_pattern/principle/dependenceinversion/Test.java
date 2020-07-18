package com.design_pattern.principle.dependenceinversion;

public class Test {

	public static void main(String[] args) {
		
		Study study = new Study();
		
		study.studyCourse(new JavaCourse());
		study.studyCourse(new PHPCourse());
		study.studyCourse(new PythonCourse());
		
	}

}
