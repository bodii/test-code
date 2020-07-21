package com.design_pattern.pattern.creational.singleton;

public enum EnumInstance {
    INSTANCE;
    private Object data;

    public void setData(Object data) {
        this.data = data;
    }

    public Object getData() {
        return data;
    }

    public static EnumInstance getInstance() {
        return INSTANCE;
    }
}