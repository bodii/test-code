package chapter07;

public class Blackhole {
    public static void enter(Object obj) {

        System.out.println("Step 1");

        magic(obj);

        System.out.println("Step 2");

        synchronized (obj) {
            System.out.println("Step 3 (never reached here)"); // 不会执行到这里
        }
    }

    public static void magic(Object obj) {
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

        try {
            synchronized (thread) {
                thread.setName("Lock");
                thread.start();
                while (thread.getName().equals("Lock")) {
                    thread.wait();
                }
            }
        } catch (InterruptedException e) {

        }

    }
}