package com.design_pattern.pattern.creational.singleton;

import java.util.HashMap;
import java.util.Map;

public class ContainerSingleton {
    private static Map<String, Object> instances = new HashMap<>();

    private ContainerSingleton() {}

    public static void putInstance(String key, Object instance) {
        if (!key.isBlank() && instance != null) {
            if (!instances.containsKey(key)) {
                instances.put(key, instance);
            }
        }
    }

    public static Object getInstance(String key) {
        return instances.get(key);
    }
}