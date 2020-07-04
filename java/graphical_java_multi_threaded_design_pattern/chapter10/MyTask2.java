package chapter10;

import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.CyclicBarrier;
import java.util.Random;

public class MyTask2 implements Runnable {
    private static final int PHASE = 5;
    private final CyclicBarrier phaseBarrier;
    private final CountDownLatch doneLatch;
    private final int context;
    private static final Random random = new Random(314159);

    public MyTask2(CyclicBarrier phaseBarrier , CountDownLatch doneLatch, int context) {
        this.phaseBarrier = phaseBarrier;
        this.doneLatch = doneLatch;
        this.context = context;
    }

    @Override
    public void run() {
        try {
            for (int phase = 0; phase < PHASE; phase ++) {
                doPhase(phase);
                phaseBarrier.await();
            }
        } catch (InterruptedException | BrokenBarrierException e) {
            e.printStackTrace();
        } finally {
            doneLatch.countDown();
        }
    }

    protected void doPhase(int phase) {
        String name = Thread.currentThread().getName();
        System.out.println(name + ":MyTask2:BEGIN:context = " + context + ", phase = " + phase);

        try {
            Thread.sleep(random.nextInt(3000));
        } catch (InterruptedException e) {

        } finally {
            System.out.println(name + "MyTask2:END:context = " + context + ", phase = " + phase);
        }
    }
}