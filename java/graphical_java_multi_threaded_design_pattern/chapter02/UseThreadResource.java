package chapter02;

import java.util.Random;

public class UseThreadResource extends Thread {
    private final static Random random = new Random(26535);
    private final BoundedResource resource;

    public UseThreadResource(BoundedResource resource) {
        this.resource = resource;
    }

    @Override
    public void run() {
        try {
            while (true) {
                resource.use();
                Thread.sleep(random.nextInt(3000));
            }
        } catch (InterruptedException e) {
            
        }
    }
 }