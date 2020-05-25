package chapter01;

import java.util.concurrent.ThreadFactory;
import java.util.concurrent.Executors;

public class TestThreadFactory {
   public static void main(String[] args) {
       ThreadFactory factory = Executors.defaultThreadFactory();
       factory.newThread(new Printer("Nice!")).start();
       for (int i = 0; i < 10000; i ++)
            System.out.print("Good!");
   } 
}