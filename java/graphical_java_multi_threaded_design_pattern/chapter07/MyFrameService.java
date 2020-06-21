package chapter07;

public class MyFrameService {
    public static void service() {
        new Thread() {
            @Override
            public void run() {
                System.out.print("service");

                for (int i = 0; i < 50; i++) {
                    System.out.print(".");
                    try {
                        Thread.sleep(100);
                    } catch (InterruptedException e) {
        
                    }
                }
        
                System.out.println("done.");
            }
        }.start();

    }
}