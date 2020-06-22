package chapter08;

import java.util.Random;

public class Request {
    public final String name;
    public final int number;
    private static final Random random = new Random();
    
    public Request(String name, int number) {
        this.name = name;
        this.number = number;
    }

    public void execute() {
        System.out.println(Thread.currentThread().getName() + " executes " + this);
        try {
            Thread.sleep(random.nextInt(1000));
        } catch (InterruptedException e) {}
    }

    @Override
    public String toString() {
        return "[ Reqeust from " + name + " No." + number + "]";
    }
}