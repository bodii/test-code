package chapter07;

public class HostTMain {
    public static void main(String[] args) {
        System.out.println("main Begin");

        HostT host = new HostT();

        host.request(10, 'A');
        host.request(20, 'B');
        host.request(30, 'C');

        System.out.println("main end");
    }
}