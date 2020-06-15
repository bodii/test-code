package chapter04;

import java.io.IOException;
import java.util.concurrent.TimeoutException;

public class HostMain {
    public static void main(String[] args) {
        Host host = new Host(10000L);
        try {
            System.out.println("execute BEGIN");
            host.execute();
        } catch (TimeoutException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}