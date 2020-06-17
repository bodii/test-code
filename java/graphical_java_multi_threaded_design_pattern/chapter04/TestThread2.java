package chapter04;

public class TestThread2 extends Thread {
    @Override
    public void run() {
        System.out.print("BEGIN");
        
        for (int i = 0; i < 50; i ++) {
            System.out.print(".");
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {

            }
        }

        System.out.println("END");
    }
}