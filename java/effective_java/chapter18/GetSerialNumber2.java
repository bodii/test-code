package chapter18;

import chapter18.GenerateSerialNumber2;

public class GetSerialNumber2 {
    public static void main(String[] args) {
        Thread thread1 = new Thread(
            () -> {
                System.out.println(
                    GenerateSerialNumber2.generateserialNumber()
                    );
            }
        );

        Thread thread2 = new Thread(
            () -> {
                int i = 10;
                while (i -- > 0)
                    System.out.println(
                        GenerateSerialNumber2.generateserialNumber()
                        );
            }
        );

        
        thread2.start();
        thread1.start();
    }
}