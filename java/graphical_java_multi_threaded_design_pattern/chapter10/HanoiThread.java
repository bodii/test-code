package chapter10;

public class HanoiThread extends Thread {
    // 发出终止请求后变为true
    private volatile boolean shutdownRequested = false;
    // 发出终止请求的时间
    private volatile long requestedTimeMillis = 0;

    // 终止请求
    public void shutdownRequest() {
        requestedTimeMillis = System.currentTimeMillis();
        shutdownRequested = true;
        interrupt();
    }

    // 检查是否发出了终止请求
    public boolean isShutdownRequested() {
        return shutdownRequested;
    }

    // 线程体
    @Override
    public void run() {
        try {
            for (int level = 0; !isShutdownRequested(); level ++) {
                System.out.println("==== Level " + level + " ==== ");
                doWork(level, 'A', 'B', 'C');
                System.out.println("");
            }
        } catch (InterruptedException e) {
            //TODO: handle exception
        } finally {
            doShutdown();
        }
    }

    // 操作
    private void doWork(int level, char A, char B, char C) 
            throws InterruptedException {
        if (level > 0) {
            doWork(level - 1, A, C, B);
            System.out.print(A + "->" + B + " ");
            doWork(level - 1, C, B, A);
        }
    }

    // 终止处理
    private void doShutdown() {
        long time = System.currentTimeMillis() - requestedTimeMillis;
        System.out.println("doShutdown: Latency = " + time + " msec.");
    }
}