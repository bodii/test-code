package com.design_pattern.pattern.creational.singleton;

public class ThreadLocalT implements Runnable {

    @Override
    public void run() {
        ThreadLocalInstance threadLocalInstance = ThreadLocalInstance.getInstance();
        System.out.println(Thread.currentThread().getName() + " instance:" + threadLocalInstance);

    }
    
}