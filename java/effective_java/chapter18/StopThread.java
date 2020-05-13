package chapter18;

import java.util.concurrent.TimeUnit;

/**
 * Broken! - How long would you expect this program to run
 * 注意 写方法(requestStop)和读方法(stopRequested)都被同步了。只同步写方法还不够！
 * 除非读和写操作都被同步，否则无法保证同步能起作用。
 */
public class StopThread {
    private static boolean stopRequested;

    private static synchronized void requestStop() {
        stopRequested = true;
    }

    private static synchronized boolean stopRequested() {
        return stopRequested;
    }

    public static void main(String[] args) throws InterruptedException {
        Thread backgroundThread = new Thread(
            () -> {
                int i = 0;
                while (!stopRequested())
                    i ++;
            }
        );
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        requestStop();
    }
}