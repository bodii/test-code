package chapter05;

import java.util.concurrent.ArrayBlockingQueue;

public class Table2 extends ArrayBlockingQueue<String> {
    public Table2(int count) {
        super(count);
    }

    @Override
    public void put(String cake) throws InterruptedException {
        System.out.println(Thread.currentThread().getName() + " puts " + cake);
        super.put(cake);
    }

    @Override
    public String take() throws InterruptedException {
        String cake = super.take();
        System.out.println(Thread.currentThread().getName() + " takes " + cake);

        return cake;
    }
}