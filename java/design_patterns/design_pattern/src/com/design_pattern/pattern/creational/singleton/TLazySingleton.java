package com.design_pattern.pattern.creational.singleton;

public class TLazySingleton implements Runnable{

    @Override
    public void run() {
        LazySingleton lazySingleton = LazySingleton.getInstance();
        System.out.println(
            Thread.currentThread().getName() + ": " + lazySingleton
            );
    }
    
}