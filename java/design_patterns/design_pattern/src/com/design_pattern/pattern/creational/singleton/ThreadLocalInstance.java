package com.design_pattern.pattern.creational.singleton;

public class ThreadLocalInstance {
        private static final ThreadLocal<ThreadLocalInstance> threadLocalInstance 
            = new ThreadLocal<>() {

                @Override
                protected ThreadLocalInstance initialValue() {
                    return new ThreadLocalInstance();
                }
            };
        
        private ThreadLocalInstance() { }

        public static ThreadLocalInstance getInstance() {
            return threadLocalInstance.get();
        }
 
        
}