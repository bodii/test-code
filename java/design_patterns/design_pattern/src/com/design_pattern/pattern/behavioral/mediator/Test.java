package com.design_pattern.pattern.behavioral.mediator;

public class Test {
    public static void main(String[] args) {
        User jack = new User("jack");
        jack.sendMessage("Hi, my name is jack");

        User tom = new User("tom");
        tom.sendMessage("Hi, my name is tom");
    }
}
