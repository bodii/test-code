package com.design_pattern.pattern.creational.singleton;

public class LazySingleton2 {
    private static LazySingleton2 lazySingleton = null;
    private static boolean flag = true;

    /**
     * 拒绝被实例化
     */
    private LazySingleton2() {
        if (flag)
            flag = false; 
        else
            throw new RuntimeException("Singleton mode prohibits reflection calls.");
    }

    public static LazySingleton2 getInstance() {
        if (lazySingleton == null) {
            lazySingleton = new LazySingleton2();
        }

        return lazySingleton;

    }
}