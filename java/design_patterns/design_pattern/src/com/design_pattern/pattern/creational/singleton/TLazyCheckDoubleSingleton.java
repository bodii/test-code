package com.design_pattern.pattern.creational.singleton;

public class TLazyCheckDoubleSingleton implements Runnable {

    @Override
    public void run() {
        LazyCheckDoubleSingleton lazyCheckDoubleSingleton = LazyCheckDoubleSingleton.getInstance();
        System.out.println(Thread.currentThread().getName() + 
            ": " + lazyCheckDoubleSingleton);
    }
    
}