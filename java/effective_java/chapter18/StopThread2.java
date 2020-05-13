package chapter18;

import java.util.concurrent.TimeUnit;

/**
 * volatile修饰不执行互斥访问，但它可以保证任何一个线程在读取该域的时候都将看到
 * 最近刚刚被写入的值
 */
public class StopThread2 {
    private static volatile boolean stopRequested;

    private static void main(String[] args) 
            throws InterruptedException {
        Thread backgroundThread = new Thread(
            () -> {
                int i = 0;
                while (!stopRequested)
                    i ++;
            }
        );

        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        stopRequested = true;
    }
}