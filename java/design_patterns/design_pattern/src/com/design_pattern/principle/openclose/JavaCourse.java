package com.design_pattern.principle.openclose;

/**
 * 
 * @author wong
 *
 */
public class JavaCourse implements ICourse{
	private Integer id;
	private String name;
	private Double price;
	
	public JavaCourse(Integer id, String name, Double price) {
		this.id = id;
		this.name = name;
		this.price = price;
	}
	
	public Integer getId() {
		return id;
	}
	
	public String getName() {
		return name;
	}
	
	public Double getPrice() {
		return price;
	}
}
