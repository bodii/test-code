package chapter04;

public class TestThreadMain2 {
    public static void main(String[] args) {
        while(true) 
            (new TestThread2()).start();
    }
}