package com.design_pattern.pattern.creational.singleton;

public class StaticInnerClassSingleton {

    private static class InnerClass {
        private static StaticInnerClassSingleton staticInnerClassSingleton = 
            new StaticInnerClassSingleton();
    } 

    public static StaticInnerClassSingleton getInstance() {
        return InnerClass.staticInnerClassSingleton;
    }

    private StaticInnerClassSingleton() {}
}