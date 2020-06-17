package chapter04;

public class TestThreadMain {
    public static void main(String[] args) {
        while(true) 
            (new Thread((new TestThread()))).start();
    }
}