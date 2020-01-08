package chapter10;

import java.util.concurrent.TimeUnit;

public class StopThread2 {
    private static boolean stopRequested;

    /**
     * 读和写操作都被同步，否则无法保证同步能起作用。
     */
    private static synchronized void requestStop() {
        stopRequested = true;
    }

    private static synchronized boolean stopRequested() {
        return stopRequested;
    }

    public static void main(String[] args) throws InterruptedException{
        Thread backgroundThread = new Thread(
            () -> {
                int i = 0;
                while(!stopRequested())
                    i ++;
            }
        );
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        requestStop();
    }
}