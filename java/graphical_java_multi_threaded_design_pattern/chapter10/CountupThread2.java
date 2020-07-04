package chapter10;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;

public class CountupThread2 extends Thread {
    // 计数值
    private long counter = 0;
    // 保存的文件路径
    private static final String savePath = "counter.txt";

    // 终止请求
    public void shutdownRequest() {
        interrupt();
    }

    // 线程体
    @Override
    public void run() {
        try {
            while (!isInterrupted()) 
                doWork();
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
        saveCounter();
    }

    private void saveCounter()  {
        try {
            File counterFile = new File(savePath);
            FileOutputStream file = new FileOutputStream(counterFile, false);
            StringBuffer data = new StringBuffer(); 
            for (int i = 1; i <= counter; i ++) 
                data.append(i + "\n");
            file.write(data.toString().getBytes());
            file.close();
        } catch (IOException e) {
            e.printStackTrace();
        }

    }
}