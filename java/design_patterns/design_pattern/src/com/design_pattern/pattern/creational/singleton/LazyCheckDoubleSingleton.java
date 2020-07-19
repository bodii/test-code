package com.design_pattern.pattern.creational.singleton;

public class LazyCheckDoubleSingleton {
    private volatile static LazyCheckDoubleSingleton lazyCheckDoubleSingleton = null;

    private LazyCheckDoubleSingleton() {}

    public static LazyCheckDoubleSingleton getInstance() {
        if (lazyCheckDoubleSingleton == null) {
            synchronized(LazyCheckDoubleSingleton.class) {
                if (lazyCheckDoubleSingleton == null)
                    lazyCheckDoubleSingleton = new LazyCheckDoubleSingleton();
            }
        }
        
        return lazyCheckDoubleSingleton;
    }
}