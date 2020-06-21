package chapter07;

import java.util.concurrent.Executor;
import java.util.concurrent.ThreadFactory;

public class LogMain {
    public static void main(String[] args) {
        Thread.currentThread().setName("MainThread");
        Log.println("main:BEGIN");
        
        new Executor(){
        
            @Override
            public void execute(Runnable r) {
                Log.println("execute:BEGIN");
                new ThreadFactory(){
                
                    @Override
                    public Thread newThread(Runnable r) {
                        Log.println("newThread:BEGIN");
                        Thread t = new Thread(r, "QuizThread");
                        Log.println("newThread:END");

                        return t;
                    }
                }.newThread(r).start();
                Log.println("execute:END");
            }
        }.execute(
            new Runnable() {
                public void run() {
                    Log.println("run:BEGIN");
                    Log.println("Hello!");
                    Log.println("run:END");
                }
            }
        );

        Log.println("main:END");
    }
}