package chapter07;

import java.util.concurrent.Executors;

public class HostMain4 {
    public static void main(String[] args) {
        System.out.println("main BEGIN");

        Host3 host = new Host3(
            Executors.defaultThreadFactory()
        );

        host.request(10, 'A');
        host.request(20, 'B');
        host.request(30, 'c');

        System.out.println("main END");
    }
}