package com.design_pattern.pattern.creational.singleton;

public class Test11 {
    public static void main(String[] args) {
        System.out.println("main Thread: " + ThreadLocalInstance.getInstance());
        System.out.println("main Thread: " + ThreadLocalInstance.getInstance());
        System.out.println("main Thread: " + ThreadLocalInstance.getInstance());
        System.out.println("main Thread: " + ThreadLocalInstance.getInstance());
        System.out.println("main Thread: " + ThreadLocalInstance.getInstance());
        System.out.println("main Thread: " + ThreadLocalInstance.getInstance());
        Thread t1 = new Thread(new ThreadLocalT());
        Thread t2 = new Thread(new ThreadLocalT());
        t1.start();
        t2.start();
        System.out.println("The program end.");
    }
}