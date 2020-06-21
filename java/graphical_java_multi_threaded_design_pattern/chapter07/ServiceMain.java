package chapter07;

import java.io.IOException;

public class ServiceMain {
    public static void main(String[] args) {
        try {
            new ServiceMini(8888).execute();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}