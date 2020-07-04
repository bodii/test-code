package chapter10;

import java.util.concurrent.CountDownLatch;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class MyTaskMain2 {
    private static final int THREADS = 3; // 线程的个数

    public static void main(String[] args) {
        System.out.println("BEGIN");

        // 由ExecutorService提供进行工作的线程
        ExecutorService service = Executors.newFixedThreadPool(THREADS);

        // 屏障被解除时的操作
        Runnable barrierAction = new Runnable() {
            @Override
            public void run() {
                System.out.println("Barrier Action!");
            }
        };

        // CyclicBarrier用于使线程步调一致
        CyclicBarrier phaseBarrier = new CyclicBarrier(THREADS, barrierAction);
        
        // CountDownLatch用于确认工作是否结束
        CountDownLatch doneLatch = new CountDownLatch(THREADS);

        try {
            // 开始工作
            for (int i = 0; i < THREADS; i++) {
                service.execute(new MyTask2(phaseBarrier, doneLatch, i));
            }
            // 等待工作结束
            System.out.println("AWAIT");
            doneLatch.await();
        } catch (InterruptedException e) {

        } finally {
            service.shutdown();
            System.out.println("END");
        }

        
    }
}