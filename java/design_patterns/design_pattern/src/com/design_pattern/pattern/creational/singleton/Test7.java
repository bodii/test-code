package com.design_pattern.pattern.creational.singleton;

import java.lang.reflect.Constructor;
import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;

public class Test7 {
    public static void main(String[] args) throws NoSuchMethodException,
            IllegalAccessException, InvocationTargetException,
            InstantiationException, NoSuchFieldException {
        LazySingleton2 instance = LazySingleton2.getInstance();

        Class objectClass = LazySingleton2.class;
        Constructor c = objectClass.getDeclaredConstructor();
        c.setAccessible(true);


        Field flag = instance.getClass().getDeclaredField("flag");
        flag.setAccessible(true);
        flag.set(instance, true);
        LazySingleton2 newInstance = (LazySingleton2) c.newInstance();

        System.out.println(instance);
        System.out.println(newInstance);
        System.out.println(instance == newInstance);
    }
}
