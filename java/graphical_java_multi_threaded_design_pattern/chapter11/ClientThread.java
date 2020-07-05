package chapter11;

public class ClientThread extends Thread {

    public ClientThread(String name) {
        super(name);
    }

    @Override
    public void run() {
        System.out.println(getName() + " BEGIN");

        for (int i = 0; i < 10; i++) {
            Log2.println("i = " + i);
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
            }
        }
        Log2.close();

        System.out.println(getName() + " END");
    }
}