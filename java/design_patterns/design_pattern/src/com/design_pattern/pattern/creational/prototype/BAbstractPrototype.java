package com.design_pattern.pattern.creational.prototype;

public class BAbstractPrototype extends AbstractPrototype {

    public static void main(String[] args) throws CloneNotSupportedException {
        AbstractPrototype a = new AbstractPrototype();
        AbstractPrototype b =  (AbstractPrototype) a.clone();
    }
    
}