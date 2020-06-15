package chapter04;

import java.util.concurrent.TimeoutException;

public class Host {
    private final long timeout; // 超时
    private boolean ready = false; // 正常执行时值为true

    public Host(Long timeout) {
        this.timeout = timeout;
    }

    // 修改状态
    public synchronized void setExecutable(boolean on) {
        ready = on;
        notifyAll();
    }

    // 检查状态之后再执行
    public synchronized void execute() throws InterruptedException, TimeoutException {
        long start = System.currentTimeMillis(); // 开始时间
        while (!ready) {
            long now = System.currentTimeMillis();
            long rest = timeout - (now - start);
            if (rest <= 0)
                throw new TimeoutException("now - start = " + (now - start) + 
                    (now -start) + ", timeout = " + timeout);

            wait(rest);
        }

        doExecute();
    }

    private void doExecute() {
        System.out.println(Thread.currentThread().getName() + 
            " calls doExecute");
    }


}