package chapter18;

import java.util.concurrent.CountDownLatch;
import java.util.concurrent.Executor;

public class TestCountdownLatch {
    public static long time(Executor executor, int concurrency, 
            Runnable action) throws InterruptedException {
        // 倒计数锁存器
        CountDownLatch ready = new CountDownLatch(concurrency);
        CountDownLatch start = new CountDownLatch(1);
        CountDownLatch done = new CountDownLatch(concurrency);

        for (int i = 0; i < concurrency; i ++) {
            executor.execute(
                () -> {
                    ready.countDown();

                    try {
                        start.await();
                        action.run();
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                    } finally {
                        done.countDown();
                    }
                }
            );
        }

        ready.await();
        long startNanos = System.nanoTime();
        start.countDown();
        done.await();


        return System.nanoTime() - startNanos;
    }
}