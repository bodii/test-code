package com.design_pattern.principle.dependenceinversion;

public class Test2 {

	public static void main(String[] args) {
		Study2 javaStudy = new Study2(new JavaCourse());
		javaStudy.studyCourse();
		
		Study2 pHPStudy2 = new Study2(new PHPCourse());
		pHPStudy2.studyCourse();
		
		Study2 pythonStudy2 = new Study2(new PythonCourse());
		pythonStudy2.studyCourse();

	}

}
