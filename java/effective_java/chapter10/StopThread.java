package chapter10;

import java.util.concurrent.TimeUnit;

public class StopThread {
    /**
     * 你可能期待这个程序运行大约一秒种左右，之后主线程将stopRequested设置为true,
     * 致使后台线程的循环终止。但是这个程序永远不会终止：因为后台线程永远在循环。
     * 
     * 问题在于，由于没有同步，就不能保证后台线程何时“看到”主线程对stopRequested
     * 的值所做的改变。
     */
    private static boolean stopResquested;

    public static void main(String[] args) throws InterruptedException {
        Thread backgroundThread = new Thread(() -> {
            int i = 0;
            while (!stopResquested)
                i++;
        });
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        stopResquested = true;
    }
}