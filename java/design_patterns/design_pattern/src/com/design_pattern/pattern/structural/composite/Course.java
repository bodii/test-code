package com.design_pattern.pattern.structural.composite;

public class Course extends CatalogComponent{
    private String name;
    private double price;

    public Course(String name, double price) {
        this.name = name;
        this.price = price;
    }

    @Override
    public void add(CatalogComponent catalogComponent) {
        super.add(catalogComponent);
    }

    @Override
    public void remove(CatalogComponent catalogComponent) {
        super.remove(catalogComponent);
    }

    @Override
    public String getName(CatalogComponent catalogComponent) {
        return name;
    }

    @Override
    public double getPrice(CatalogComponent catalogComponent) {
        return price;
    }

    @Override
    public void print() {
        System.out.println("Course Name: " + name + " Price:" + price);
    }
}
