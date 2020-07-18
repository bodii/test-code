package com.design_pattern.principle.openclose;

/**
 * 
 * @author wong
 *
 */
public class JavaDisCountCourse extends JavaCourse {

	public JavaDisCountCourse(Integer id, String name, Double price) {
		super(id, name, price);
	}
	
	public Double getOrginPrice() {
		return super.getPrice();
	}
	
	@Override
	public Double getPrice() {
		return super.getPrice() * 0.8;
	}
	
}
