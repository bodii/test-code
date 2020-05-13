package chapter18;

import java.io.PrintStream;

import chapter18.GenerateSerialNumber;

public class GetSerialNumber {
    public static void main(String[] args) {
        Thread thread1 = new Thread(
            () -> {
                System.out.println(
                    GenerateSerialNumber.generateSerialNumber()
                    );
            }
        );

        Thread thread2 = new Thread(
            () -> {
                int i = 10;
                while (i -- > 0)
                    System.out.println(
                        GenerateSerialNumber.generateSerialNumber()
                        );
            }
        );

        thread1.start();
        thread2.start();
    }
}