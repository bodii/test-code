package chapter01;

public class TestPrintThread {
    public static void main(String[] args) {
        new PrintThread("Good!").start();
        new PrintThread("Nice!").start();
    }
}