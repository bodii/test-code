package com.design_pattern.pattern.creational.singleton;

public class Test {
    public static void main(String[] args) {
        LazySingleton lazySingleton = LazySingleton.getInstance();
        System.out.println(lazySingleton);

        LazySingleton lazySingleton2 = LazySingleton.getInstance();
        System.out.println(lazySingleton2);
    }
}