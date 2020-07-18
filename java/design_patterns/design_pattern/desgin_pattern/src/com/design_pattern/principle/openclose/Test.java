package com.design_pattern.principle.openclose;

/**
 * 
 * @author wong
 *
 */
public class Test {

	public static void main(String[] args) {
//		ICourse javaCourse = new JavaCourse(235, "java", 38d);
//		System.out.println(
//				"id: " + javaCourse.getId() 
//				+ " 课程: " + javaCourse.getName() 
//				+ " 学费： " + javaCourse.getPrice()
//				);
		
		JavaDisCountCourse javaDiscountCourse  = new JavaDisCountCourse(235, "java", 38d);
		System.out.println(
				"id:" + javaDiscountCourse.getId()
				+ " 课程:" + javaDiscountCourse.getName()
				+ " 原价:" + javaDiscountCourse.getOrginPrice()
				+ " 折扣价:" + javaDiscountCourse.getPrice()
				);
	}

}
