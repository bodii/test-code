package chapter07;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class HostMain6 {
    public static void main(String[] args) {
        System.out.println("main BEGIN");

        ExecutorService executorservice = Executors.newCachedThreadPool();
        Host5 host = new Host5(
            executorservice
        );

        try {
            host.request(10, 'A');
            host.request(20, 'b');
            host.request(30, 'c');
        } finally {
            executorservice.shutdown();  
        }

        System.out.println("main END");
    }
}