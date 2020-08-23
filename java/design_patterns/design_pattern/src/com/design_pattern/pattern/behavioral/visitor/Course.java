package com.design_pattern.pattern.behavioral.visitor;

public abstract class Course {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public abstract void accept(Ivisitor visitor);
}
