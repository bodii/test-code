package com.design_pattern.pattern.creational.singleton;

public class LazySingleton {
    private static LazySingleton lazySingleton = null;

    /**
     * 拒绝被实例化
     */
    private LazySingleton() {}

    public static LazySingleton getInstance() {
        if (lazySingleton == null) {
            lazySingleton = new LazySingleton();
        }

        return lazySingleton;

    }
}