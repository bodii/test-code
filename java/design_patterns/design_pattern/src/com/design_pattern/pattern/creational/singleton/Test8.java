package com.design_pattern.pattern.creational.singleton;

import java.lang.reflect.Constructor;
import java.lang.reflect.InvocationTargetException;

public class Test8 {
    public static void main(String[] args) throws NoSuchMethodException, IllegalAccessException, InvocationTargetException, InstantiationException {
        LazySingleton3 instance = LazySingleton3.getInstance();
        Class objectClass = LazySingleton3.class;
        Constructor c = objectClass.getDeclaredConstructor();
        c.setAccessible(true);
        LazySingleton3 newInstance = (LazySingleton3) c.newInstance();

        System.out.println("instance = " + instance);
        System.out.println("newInstance = " + newInstance);
        System.out.println(instance == newInstance);
    }
}
