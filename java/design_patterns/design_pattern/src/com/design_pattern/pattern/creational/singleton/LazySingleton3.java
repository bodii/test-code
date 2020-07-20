package com.design_pattern.pattern.creational.singleton;

public class LazySingleton3 {
    private static LazySingleton3 lazySingleton = null;

    /**
     * 拒绝被实例化
     */
    private LazySingleton3() {
        if (lazySingleton != null)
            throw new RuntimeException("Singleton mode prohibits reflection calls.");
    }

    public static LazySingleton3 getInstance() {
        if (lazySingleton == null) {
            lazySingleton = new LazySingleton3();
        }

        return lazySingleton;

    }
}