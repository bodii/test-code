package com.design_pattern.principle.singleresponsibility;

public class UserInfo {
    private String name;
    private int age;

    public void updateUserInfo(String name, int age) {
        this.name = name;
        this.age = age;
    }

    public void updateUserInfo(String name, String... properties ) {
        this.name = name;
    }

    public void updateUsername(String name) {
        this.name = name;
    }

    public void updateUserAge(int age) {
        this.age = age;
    }
}