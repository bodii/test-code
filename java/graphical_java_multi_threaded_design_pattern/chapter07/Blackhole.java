package chapter07;

public class Blackhole {
    public static void enter(Object obj) {

        System.out.println("Step 1");
        try {
        magic(obj);
        } catch (InterruptedException e) {

        }

        System.out.println("Step 2");

        synchronized (obj) {
            System.out.println("Step 3 (never reached here)"); // 不会执行到这里
        }
    }

    public static void magic(Object obj) throws InterruptedException {
        Thread thread = new Thread() {
            @Override
            public void run() {
                synchronized (obj) {
                    synchronized (this) {
                        this.setName("LockNow");
                        this.notifyAll();
                    }

                    while(true) {

                    }
                }
            }
        };

        synchronized (thread) {
            thread.setName("Lock");
            thread.start();
            while (thread.getName().equals("Lock")) {
                thread.wait();
            }
        }


    }
}