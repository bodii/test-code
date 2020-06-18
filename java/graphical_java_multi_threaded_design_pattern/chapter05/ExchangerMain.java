package chapter05;

import java.util.concurrent.Exchanger;

public class ExchangerMain {
    public static void main(String[] args) {
        Exchanger<char[]> exchanger = new Exchanger<>();
        char[] buffer1 = new char[10];
        char[] buffer2 = new char[10];
        new ProducterThread(exchanger, buffer1, 314159).start();
        new ConsumerThread(exchanger, buffer2, 265358).start();            
    }
}