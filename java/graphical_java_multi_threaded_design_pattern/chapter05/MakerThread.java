package chapter05;

import java.util.Random;

public class MakerThread extends Thread {

    private final Random random;
    private final Table table;
    private static int id = 0; // 蛋糕的流水号

    public MakerThread(String name, Table table, long seed) {
        super(name);
        this.table = table;
        random = new Random(seed);
    }

    @Override 
    public void run () {
        try {
            while (true) {
                Thread.sleep(random.nextInt(1000));
                String cake = "[ cake No." + nextId() + " by " + getName() + " ]";
                table.put(cake);
            }
        } catch (InterruptedException e) {
            
        }
    }

    private static synchronized int nextId () {
        return id++;
    }
}