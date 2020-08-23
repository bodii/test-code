package com.design_pattern.pattern.behavioral.visitor;

public class FreeCourse extends Course {
    @Override
    public void accept(Ivisitor visitor) {
        visitor.visit(this);
    }
}
