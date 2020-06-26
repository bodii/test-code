package chapter08;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.RejectedExecutionException;
import java.util.Random;

public class ClientThread2 extends Thread {
    private final ExecutorService executorService;
    private static final Random random = new Random();

    public ClientThread2(String name, ExecutorService executorService) {
        super(name);
        this.executorService = executorService;
    }

    @Override
    public void run() {
        try {
            for (int i = 0; true; i ++) {
                Request2 request = new Request2(getName(), i);
                executorService.execute(request);

                Thread.sleep(random.nextInt(1000));
            }
        } catch (InterruptedException e) {

        } catch (RejectedExecutionException e) {
            System.out.println(getName() + " : " + e);
        }
    }
}