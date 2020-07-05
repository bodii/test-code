package chapter11;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class Log2Main2 {
    private static final int TASKS = 10;

    public static void main(String[] args) {
        // 要运行的服务
        ExecutorService service = Executors.newFixedThreadPool(3);
        try {
            // 写日志的任务
            Runnable printTask = new Runnable() {
                @Override
                public void run() {
                    for (int i = 0; i < TASKS; i++) {
                        Log2.println("Hello!");
                    }
                    Log2.close();
                }
            };
            for (int i = 0; i < 3; i++) {
                // 执行任务
                service.execute(printTask);              
            }
        } finally {
            service.shutdown();
        }
    }
}