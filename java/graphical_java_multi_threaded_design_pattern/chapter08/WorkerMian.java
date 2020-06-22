package chapter08;

public class WorkerMian {
    public static void main(String[] args) {
        Channel channel = new Channel(5); 
        channel.startWorkers();
        new ClientThread("Alice", channel).start();
        new ClientThread("Bobby", channel).start();
        new ClientThread("Chris", channel).start();
    }
}