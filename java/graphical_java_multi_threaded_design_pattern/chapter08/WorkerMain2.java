package chapter08;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class WorkerMain2 {
    public static void main(String[] args) {
        ExecutorService executorService = Executors.newFixedThreadPool(5);

        try {
            new ClientThread2("Alice", executorService).start();
            new ClientThread2("Bobby", executorService).start();
            new ClientThread2("Chris", executorService).start();

            // 等待大约5秒
            Thread.sleep(5000);
        } catch (InterruptedException e) {
        } finally {
            executorService.shutdown();
        }
    }
}