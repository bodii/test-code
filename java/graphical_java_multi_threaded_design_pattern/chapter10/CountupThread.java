package chapter10;

public class CountupThread extends Thread {
    // 计数值
    private long counter = 0;

    // 发出终止请求后变为true
    private volatile boolean shutdownRequested = false;

    //  终止请求
    public void shutdownRequest() {
        shutdownRequested = true;
        interrupt();
        /*
        只将shutdownRequested标志设为true是不行的
        因为当想要终止线程时，该线程可能正在sleep、wait。而当线程正在sleep、wait时，
        即使将shudownRequested标志为true,线程也不会从等待队列中立刻出来，所以我们必
        须使用interrupt方法对线程下达"中断"的指示。
        如果调用interrupt方法后，如果线程正在sleep或是wait，那么会抛出InterruptedException
        异常，而如果不抛出异常，线程就会变为中断状态。也就是说，没有必要特意准备一个新的shutdownRequested
        标志。只要捕获InterruptedException,使用isInterrupted方法来检查线程是否处于中断状态
        但是，只要线程正在执行的方法中有一处忽略InterruptedException，上面的方法就可能行不通。
        try {
            Thread.sleep(100);
        } catch (InterrruptedException e) {

        }
        这样，即使在wait、sleep、join状态时抛出了InterruptedException，线程也不会变为中断状态。
        也就是说，如果程序中没有shutdownRequested标志，而且有上面这样的代码，那么即使使用shutdownRequest
        方法发出了终止请求，该请求也不会被处理。

        */
    }

    // 检查是否发出了终止请求
    public boolean isShutdownReqeusted() {
        return shutdownRequested;
    }

    // 线程体
    @Override
    public final void run() {
        try {
            while (!isShutdownReqeusted()) {
                doWork();
            }
        } catch (InterruptedException e) {

        } finally {
            doShutdown();
        }
    }

    // 操作
    private void doWork() throws InterruptedException {
        counter ++;
        System.out.println("doWork: counter = " + counter);
        Thread.sleep(500);
    }

    // 终止处理
    private void doShutdown() {
        System.out.println("doShutdown: counter = " + counter);
    }
}