package com.design_pattern.pattern.behavioral.visitor;

public class CodingCourse extends Course {
    private int price;

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    @Override
    public void accept(Ivisitor visitor) {
        visitor.visit(this);
    }
}
