package chapter12;

import java.util.concurrent.CancellationException;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;
import java.util.concurrent.RejectedExecutionException;

public class MakerClientThread2 extends Thread {
    private final ActiveObject2 activeObject;
    private final char fillchar;

    public MakerClientThread2(String name, ActiveObject2 activeObject) {
        super(name);
        this.activeObject = activeObject;
        this.fillchar = name.charAt(0);
    }

    @Override
    public void run() {
        try {
            for (int i = 0; true; i ++) {
                // 有返回值的调用
                Future<String> future = activeObject.makeString(i, fillchar);
                Thread.sleep(10);
                String value = future.get();
                System.out.println(Thread.currentThread().getName() + ": value = " + value);
            }
        } catch (
                RejectedExecutionException | CancellationException | ExecutionException
                | InterruptedException e
                ){
            System.out.println(Thread.currentThread().getName() + ":" + e);
        }
    }
}