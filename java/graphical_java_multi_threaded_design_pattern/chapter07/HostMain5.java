package chapter07;

import java.util.concurrent.Executor;

public class HostMain5 {
    public static void main(String[] args) {
        System.out.println("main BEGIN");

        Host5 host = new Host5(
            new Executor(){
            
                @Override
                public void execute(Runnable r) {
                    new Thread(r).start();
                }
            }
        );

        host.request(10, 'A');
        host.request(20, 'B');
        host.request(30, 'c');

        System.out.println("main END");
    }
}