package chapter10;

import java.util.concurrent.TimeUnit;

/**
 * volatile
 *  它更加简洁，性能也可能更好。
 * 虽然volatile修饰符不执行互斥访问，但它可以保证任何一个线程在读取该域的时
 * 候都将看到最近刚刚被写入的值
 */
public class StopThread3 {
    private static volatile boolean stopRequested;

    public static void main(String[] args) throws InterruptedException {
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