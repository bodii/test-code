package chapter07;

public class BlackholeMain {
    public static void main(String[] args) {
        System.out.println("BEGIN");

        Object obj = new Object();
        Blackhole.enter(obj);
        
        System.out.println("END");
    }
}